package emitter

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

type Disassembler struct {
	offset int

	ch *Chunk
	w  *bufio.Writer
}

func NewDisassembler(w io.Writer, ch *Chunk) *Disassembler {
	return &Disassembler{
		ch: ch,
		w:  bufio.NewWriter(w),
	}
}

func (d *Disassembler) Disassemble() {
	for d.offset = 0; d.offset < len(d.ch.code); {
		d.disassembleInstruction()
	}

	_ = d.w.Flush()
}

func (d *Disassembler) move(n int) {
	if d.offset+n > len(d.ch.code) {
		panic("offset out of bounds")
	}

	d.offset += n
}

func (d *Disassembler) disassembleInstruction() {
	opcode := d.readByte()

	switch opcode {
	case LoadConst:
		d.loadConstInstruction()

	case FnDecl:
		d.funDeclInstruction()

	case DefineSymbol:
		d.defineSymbolInstruction()

	case GetSymbol:
		d.getSymbolInstruction()

	case Call:
		d.callInstruction()

	case BeginScope:
		d.simpleInstruction("BEGINSCOPE")

	case EndScope:
		d.simpleInstruction("ENDSCOPE")

	case Pop:
		d.simpleInstruction("POP")

	case Return:
		d.simpleInstruction("RETURN")

	}
}

func (d *Disassembler) simpleInstruction(name string) {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-10s\n", name)
}

func (d *Disassembler) loadConstInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-10s", "LOADCONST")

	constIdx := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " #%d\n", constIdx)
}

func (d *Disassembler) callInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-10s", "CALL")

	argCount := d.readByte()

	_, _ = fmt.Fprintf(d.w, " args=%d\n", argCount)
}

func (d *Disassembler) funDeclInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-10s", "FNDECL")

	fnNameIdx := d.readUint32()
	entryPoint := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " name=#%d entry=%d\n", fnNameIdx, entryPoint)
}

func (d *Disassembler) defineSymbolInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-10s", "DEFSYMBOL")

	nameIdx := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " name=#%d\n", nameIdx)
}

func (d *Disassembler) getSymbolInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-10s", "GETSYMBOL")

	nameIdx := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " name=#%d\n", nameIdx)
}

func (d *Disassembler) readByte() byte {
	b := d.ch.code[d.offset]
	d.move(1)
	return b
}

func (d *Disassembler) readUint32() uint32 {
	buf := d.ch.code[d.offset : d.offset+4]
	b := binary.BigEndian.Uint32(buf)
	d.move(4)
	return b
}
