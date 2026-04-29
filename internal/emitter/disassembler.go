package emitter

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
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
		log.Panic("offset out of bounds")
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

	case JumpTo:
		d.jumpToInstruction()

	case JumpIfFalse:
		d.jumpIfFalseInstruction()

	case Call:
		d.callInstruction()

	case BeginScope:
		d.simpleInstruction("BEGIN_SCOPE")

	case EndScope:
		d.simpleInstruction("END_SCOPE")

	case Pop:
		d.simpleInstruction("POP")

	case Return:
		d.simpleInstruction("RETURN")

	}
}

func (d *Disassembler) simpleInstruction(name string) {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s\n", name)
}

func (d *Disassembler) loadConstInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s", "LOAD_CONST")

	constIdx := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " #%d\n", constIdx)
}

func (d *Disassembler) callInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s", "CALL")

	argCount := d.readByte()

	_, _ = fmt.Fprintf(d.w, " args=%d\n", argCount)
}

func (d *Disassembler) funDeclInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s", "FN_DECL")

	fnNameIdx := d.readUint32()
	entryPoint := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " name=#%d entry=%d\n", fnNameIdx, entryPoint)
}

func (d *Disassembler) defineSymbolInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s", "DEF_SYMBOL")

	nameIdx := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " name=#%d\n", nameIdx)
}

func (d *Disassembler) getSymbolInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s", "GET_SYMBOL")

	nameIdx := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " name=#%d\n", nameIdx)
}

func (d *Disassembler) jumpToInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s", "JUMP_TO")

	jumpTo := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " to=%d\n", jumpTo)
}

func (d *Disassembler) jumpIfFalseInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-12s", "JUMP_IF_FALSE")

	jumpTo := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " to=%d\n", jumpTo)
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

// TODO: I want to format the disassembled output in a more human-readable way, maybe something like this:
/*

ADDR  OPCODE      OPERANDS            HUMAN READABLE
------------------------------------------------------------
000   FN_DECL     name=#0, entry=9    ; function main()
009   BEGIN_SCOPE
010     LOAD_CONST  #1                ; push 10
015     DEF_SYMBOL  #2                ; var x = stack.pop()
...

*/
