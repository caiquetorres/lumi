package constpool

import (
	"bytes"
	"encoding/binary"
)

type typeByte byte

const (
	typeString typeByte = iota + 1
)

func (c *ConstantPool) Serialize() []byte {
	var buf bytes.Buffer
	for _, co := range c.constants {
		c.serializeConstant(co, &buf)
	}

	return buf.Bytes()
}

func (c *ConstantPool) serializeConstant(val any, buf *bytes.Buffer) {
	switch val := val.(type) {
	case string:
		c.serializeString(val, buf)

	default:
		panic("unsupported constant type")
	}
}

func (c *ConstantPool) serializeString(value string, buf *bytes.Buffer) {
	_ = buf.WriteByte(byte(typeString))

	strBytes := []byte(value)
	strLen := len(strBytes)

	var lenBuf [4]byte
	binary.BigEndian.PutUint32(lenBuf[:], uint32(strLen))

	_, _ = buf.Write(lenBuf[:])
	_, _ = buf.WriteString(value)
}
