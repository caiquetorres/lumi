package vm

import (
	"encoding/binary"
	"io"
)

type cursor struct {
	pc   uint32
	data []byte
}

func newCursor(data []byte) *cursor {
	return &cursor{
		pc:   0,
		data: data,
	}
}

func (c *cursor) moveTo(offset uint32) {
	c.pc = offset
}

func (c *cursor) hasMore() bool {
	return c.pc < uint32(len(c.data))
}

func (c *cursor) readUint8() (uint8, error) {
	const uint8Size = 1

	if c.pc+uint8Size > uint32(len(c.data)) {
		return 0, io.EOF
	}

	value := c.data[c.pc]
	c.pc += uint8Size

	return value, nil
}

func (c *cursor) readUint32() (uint32, error) {
	const uint32Size = 4

	if c.pc+uint32Size > uint32(len(c.data)) {
		return 0, io.EOF
	}

	value := binary.BigEndian.Uint32(c.data[c.pc : c.pc+uint32Size])

	c.pc += uint32Size

	return value, nil
}
