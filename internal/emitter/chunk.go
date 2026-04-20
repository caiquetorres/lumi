package emitter

import (
	"encoding/binary"
	"fmt"
)

const defaultChunkSize = 512

type Chunk struct {
	ip uint32

	code []byte
	pool *constantPool
}

func newChunk() *Chunk {
	return &Chunk{
		ip:   0,
		code: make([]byte, 0, defaultChunkSize),
		pool: newConstantPool(),
	}
}

func (c *Chunk) emit(b byte) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, b)
	c.ip++
	return
}

func (c *Chunk) patch(offset uint32, b byte) {
	if offset < uint32(len(c.code)) {
		c.code[offset] = b
	}
}

func (c *Chunk) patchUint32(offset uint32, value uint32) {
	if offset+4 <= uint32(len(c.code)) {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], value)
		copy(c.code[offset:offset+4], buf[:])
	}
}

func (c *Chunk) emitBytes(bytes ...byte) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, bytes...)
	c.ip += uint32(len(bytes))
	return
}

func (c *Chunk) emitUint8(value uint8) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, value)
	c.ip++
	return
}

func (c *Chunk) emitUint16(value uint16) (offset uint32) {
	offset = c.ip
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], value)

	c.code = append(c.code, buf[:]...)
	c.ip += 2
	return
}

func (c *Chunk) emitUint24(value uint32) (offset uint32) {
	offset = c.ip
	var buf [3]byte
	buf[0] = byte((value >> 16) & 0xFF)
	buf[1] = byte((value >> 8) & 0xFF)
	buf[2] = byte(value & 0xFF)

	c.code = append(c.code, buf[:]...)
	c.ip += 3
	return
}

func (c *Chunk) emitUint32(value uint32) (offset uint32) {
	offset = c.ip
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], value)

	c.code = append(c.code, buf[:]...)
	c.ip += 4
	return
}

func (c *Chunk) Disassemble() {
	for offset := 0; offset < len(c.code); {
		offset = c.disassembleInstruction(offset)
	}
}

func (c *Chunk) disassembleInstruction(offset int) int {
	opcode := c.code[offset]

	switch opcode {
	case FnDecl:
		return c.funDeclInstruction("FN_DECL", offset)
	case DefineSymbol:
		return c.varDeclInstruction("VAR_DECL", offset)
	case LoadConst:
		return c.constantInstruction("CONSTANT", offset)
	case Return:
		return c.simpleInstruction("RETURN", offset)
	default:
		return offset + 1
	}
}

func (c *Chunk) varDeclInstruction(name string, offset int) int {
	if offset+5 > len(c.code) {
		fmt.Printf("% 4d %s <truncated>\n", offset, name)
		return len(c.code)
	}

	nameIdx := binary.BigEndian.Uint32(c.code[offset+1 : offset+5])
	fmt.Printf("% 4d %s name=%d\n", offset, name, nameIdx)

	return offset + 5
}

func (c *Chunk) funDeclInstruction(name string, offset int) int {
	if offset+10 > len(c.code) {
		fmt.Printf("% 4d %s <truncated>\n", offset, name)
		return len(c.code)
	}

	fnNameIdx := binary.BigEndian.Uint32(c.code[offset+1 : offset+5])
	paramCount := c.code[offset+5]

	paramsOffset := offset + 6
	entryPointOffset := paramsOffset + int(paramCount)*4

	if entryPointOffset+4 > len(c.code) {
		fmt.Printf("% 4d %s <truncated>\n", offset, name)
		return len(c.code)
	}

	params := make([]uint32, 0, paramCount)
	for i := 0; i < int(paramCount); i++ {
		start := paramsOffset + i*4
		paramIdx := binary.BigEndian.Uint32(c.code[start : start+4])
		params = append(params, paramIdx)
	}

	entryPoint := binary.BigEndian.Uint32(c.code[entryPointOffset : entryPointOffset+4])

	fmt.Printf(
		"% 4d %s name=%d arity=%d params=%v entry=%d\n",
		offset,
		name,
		fnNameIdx,
		paramCount,
		params,
		entryPoint,
	)

	return entryPointOffset + 4
}

func (c *Chunk) simpleInstruction(name string, offset int) int {
	fmt.Printf("% 4d %s\n", offset, name)
	return offset + 1
}

func (c *Chunk) constantInstruction(name string, offset int) int {
	constant := c.code[offset+1]
	fmt.Printf("% 4d %s value=%d\n", offset, name, constant)
	return offset + 2
}
