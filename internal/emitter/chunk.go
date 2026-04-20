package emitter

import "encoding/binary"

const defaultChunkSize = 512

type chunk struct {
	ip uint32

	code []byte
	pool *constantPool
}

func newChunk() *chunk {
	return &chunk{
		ip:   0,
		code: make([]byte, 0, defaultChunkSize),
		pool: newConstantPool(),
	}
}

func (c *chunk) emit(b byte) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, b)
	c.ip++
	return
}

func (c *chunk) patch(offset uint32, b byte) {
	if offset < uint32(len(c.code)) {
		c.code[offset] = b
	}
}

func (c *chunk) patchUint32(offset uint32, value uint32) {
	if offset+4 <= uint32(len(c.code)) {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], value)
		copy(c.code[offset:offset+4], buf[:])
	}
}

func (c *chunk) emitBytes(bytes ...byte) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, bytes...)
	c.ip += uint32(len(bytes))
	return
}

func (c *chunk) emitUint8(value uint8) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, value)
	c.ip++
	return
}

func (c *chunk) emitUint16(value uint16) (offset uint32) {
	offset = c.ip
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], value)

	c.code = append(c.code, buf[:]...)
	c.ip += 2
	return
}

func (c *chunk) emitUint24(value uint32) (offset uint32) {
	offset = c.ip
	var buf [3]byte
	buf[0] = byte((value >> 16) & 0xFF)
	buf[1] = byte((value >> 8) & 0xFF)
	buf[2] = byte(value & 0xFF)

	c.code = append(c.code, buf[:]...)
	c.ip += 3
	return
}

func (c *chunk) emitUint32(value uint32) (offset uint32) {
	offset = c.ip
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], value)

	c.code = append(c.code, buf[:]...)
	c.ip += 4
	return
}

// func (c *chunk) disassemble() {
// 	for offset := 0; offset < len(c.code); {
// 		offset = c.disassembleInstruction(offset)
// 	}
// }

// func (c *chunk) disassembleInstruction(offset int) int {
// 	opcode := c.code[offset]

// 	switch opcode {
// 	case Return:
// 		return c.simpleInstruction("RETURN", offset)
// 	case LoadConst:
// 		return c.constantInstruction("CONSTANT", offset)
// 	}

// 	return 0
// }

// func (c *chunk) simpleInstruction(name string, offset int) int {
// 	fmt.Printf("%04d %s\n", offset, name)
// 	return offset + 1
// }

// func (c *chunk) constantInstruction(name string, offset int) int {
// 	constant := c.code[offset+1]
// 	fmt.Printf("%04d %s %d\n", offset, name, constant)
// 	return offset + 2
// }
