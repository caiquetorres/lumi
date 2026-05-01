package constpool

import (
	"bytes"
	"encoding/binary"
	"log"
)

const (
	typeBool byte = iota + 1
	typeString
	typeInt
)

func (c *ConstantPool) Serialize() []byte {
	var buf bytes.Buffer
	for _, co := range c.constants {
		c.serializeConstant(co, &buf)
	}

	return buf.Bytes()
}

func (c *ConstantPool) serializeConstant(val any, buf *bytes.Buffer) uint32 {
	switch val := val.(type) {
	case bool:
		c.serializeBool(val, buf)

	case string:
		c.serializeString(val, buf)

	case int:
		c.serializeInt(val, buf)

	default:
		log.Panic("unsupported constant type")
	}

	return 0
}

func (c *ConstantPool) serializeBool(value bool, buf *bytes.Buffer) {
	_ = buf.WriteByte(typeBool)

	if value {
		_ = buf.WriteByte(1)
	} else {
		_ = buf.WriteByte(0)
	}
}

func (c *ConstantPool) serializeString(value string, buf *bytes.Buffer) {
	_ = buf.WriteByte(typeString)
	strBytes := []byte(value)
	strLen := uint32(len(strBytes))

	var lenBuf [4]byte
	binary.BigEndian.PutUint32(lenBuf[:], strLen)

	{
		// TODO: Use varint encoding for string lengths to save space for short strings.
		_, _ = buf.Write(lenBuf[:])
	}

	_, _ = buf.WriteString(value)
}

func (c *ConstantPool) serializeInt(value int, buf *bytes.Buffer) {
	_ = buf.WriteByte(typeInt)

	var intBuf [8]byte
	binary.BigEndian.PutUint64(intBuf[:], uint64(value))

	_, _ = buf.Write(intBuf[:])
}
