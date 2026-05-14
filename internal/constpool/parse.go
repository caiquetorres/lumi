package constpool

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func Parse(data []byte) (*ConstantPool, error) {
	c := New()
	if len(data) == 0 {
		return c, nil
	}

	b := bufio.NewReader(bytes.NewReader(data))
	for {
		ty, err := b.ReadByte()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		_, err = c.internConstantFromType(typeByte(ty), b)
		if err != nil {
			return nil, fmt.Errorf("failed to parse constant: %w", err)
		}
	}

	return c, nil
}

func (c *ConstantPool) internConstantFromType(ty typeByte, b *bufio.Reader) (uint32, error) {
	switch ty {
	case typeString:
		return c.internString(b)

	default:
		return 0, fmt.Errorf("constant type %d", ty)
	}
}

func (c *ConstantPool) internString(b *bufio.Reader) (uint32, error) {
	var lenBuf [4]byte
	if _, err := b.Read(lenBuf[:]); err != nil {
		return 0, fmt.Errorf("failed to read string length: %w", err)
	}

	strLen := binary.BigEndian.Uint32(lenBuf[:])
	strBytes := make([]byte, strLen)

	_, err := b.Read(strBytes)
	if err != nil {
		return 0, fmt.Errorf("failed to read string constant: %w", err)
	}

	value := string(strBytes)
	return c.InternConstant(value), nil
}
