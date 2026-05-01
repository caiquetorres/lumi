package constpool

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func ParseConstantPool(data []byte) (*ConstantPool, error) {
	c := New()
	if len(data) == 0 {
		return c, nil
	}

	b := bufio.NewReader(bytes.NewReader(data))
	for {
		typeByte, err := b.ReadByte()
		if err != nil {
			break
		}

		if _, err := c.internConstantFromType(typeByte, b); err != nil {
			return nil, fmt.Errorf("failed to parse constant: %w", err)
		}
	}

	return c, nil
}

func (c *ConstantPool) internConstantFromType(typeByte byte, b *bufio.Reader) (uint32, error) {
	switch typeByte {
	case typeBool:
		return c.internBool(b)

	case typeString:
		return c.internString(b)

	case typeInt:
		return c.internInt(b)

	default:
		return 0, fmt.Errorf("constant type %d", typeByte)
	}
}

func (c *ConstantPool) internBool(b *bufio.Reader) (uint32, error) {
	valueByte, err := b.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("failed to read bool constant: %w", err)
	}

	value := valueByte != 0
	return c.InternConstant(value), nil
}

func (c *ConstantPool) internString(b *bufio.Reader) (uint32, error) {
	var lenBuf [4]byte
	if _, err := b.Read(lenBuf[:]); err != nil {
		return 0, fmt.Errorf("failed to read string length: %w", err)
	}

	strLen := binary.BigEndian.Uint32(lenBuf[:])

	strBytes := make([]byte, strLen)
	if _, err := b.Read(strBytes); err != nil {
		return 0, fmt.Errorf("failed to read string constant: %w", err)
	}

	value := string(strBytes)
	return c.InternConstant(value), nil
}

func (c *ConstantPool) internInt(b *bufio.Reader) (uint32, error) {
	var intBuf [8]byte
	if _, err := b.Read(intBuf[:]); err != nil {
		return 0, fmt.Errorf("failed to read int constant: %w", err)
	}

	value := int(binary.BigEndian.Uint64(intBuf[:]))
	return c.InternConstant(value), nil
}
