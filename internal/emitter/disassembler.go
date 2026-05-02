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
	case PushString:
		d.pushStringInstruction()

	case PushInt:
		d.uint32Instruction("PUSH_INT")

	case PushTrue:
		d.simpleInstruction("PUSH_TRUE")

	case PushFalse:
		d.simpleInstruction("PUSH_FALSE")

	case PushFn:
		d.pushFunctionInstruction()

	case PushNativeFn:
		d.uint32Instruction("PUSH_NATIVE_FN")

	case StoreLocal:
		d.uint32Instruction("STORE_LOCAL")

	case LoadLocal:
		d.uint32Instruction("LOAD_LOCAL")

	case JumpTo:
		d.uint32Instruction("JUMP_TO")

	case JumpIfFalse:
		d.uint32Instruction("JUMP_IF_FALSE")

	case Call:
		d.callInstruction()

	case Pop:
		d.simpleInstruction("POP")

	case Return:
		d.simpleInstruction("RETURN")

	case Add:
		d.simpleInstruction("ADD")

	case Sub:
		d.simpleInstruction("SUB")

	case Mul:
		d.simpleInstruction("MUL")

	case Div:
		d.simpleInstruction("DIV")

	case Eq:
		d.simpleInstruction("EQ")

	case Less:
		d.simpleInstruction("LESS")

	case LessEq:
		d.simpleInstruction("LESS_EQ")

	case Not:
		d.simpleInstruction("NOT")
	}
}

func (d *Disassembler) simpleInstruction(name string) {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-16s\n", name)
}

func (d *Disassembler) uint32Instruction(name string) {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-16s", name)

	value := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " %d\n", value)
}

func (d *Disassembler) pushStringInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-16s", "PUSH_STRING")

	constIdx := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " #%d\n", constIdx)
}

func (d *Disassembler) pushFunctionInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-16s", "PUSH_FN")

	fnID := d.readUint32()
	fnAddr := d.ch.fnTable[fnID]

	_, _ = fmt.Fprintf(d.w, " addr=%d", fnAddr)
	_, _ = fmt.Fprintf(d.w, "    %d\n", fnID)
}

func (d *Disassembler) callInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-16s", "CALL")

	argCount := d.readByte()

	_, _ = fmt.Fprintf(d.w, " args=%d\n", argCount)
}

func (d *Disassembler) funDeclInstruction() {
	_, _ = fmt.Fprintf(d.w, "% 4d ", d.offset-1)
	_, _ = fmt.Fprintf(d.w, "%-16s", "FN_DECL")

	fnNameIdx := d.readUint32()
	entryPoint := d.readUint32()

	_, _ = fmt.Fprintf(d.w, " name=#%d entry=%d\n", fnNameIdx, entryPoint)
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
