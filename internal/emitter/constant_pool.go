package emitter

import (
	"bytes"
	"fmt"
)

type constantPool struct {
	constants  []any
	indexByKey map[string]uint32
}

func newConstantPool() *constantPool {
	return &constantPool{
		constants:  make([]any, 0),
		indexByKey: make(map[string]uint32),
	}
}

func (c *constantPool) internConstant(value any) uint32 {
	if c.indexByKey == nil {
		c.indexByKey = make(map[string]uint32)
	}

	key := constantKey(value)
	if idx, ok := c.indexByKey[key]; ok {
		return idx
	}

	c.constants = append(c.constants, value)
	idx := uint32(len(c.constants) - 1)
	c.indexByKey[key] = idx

	return idx
}

func (c *constantPool) serialize() []byte {
	var buf bytes.Buffer

	for i, co := range c.constants {
		fmt.Fprintf(&buf, "#%d: %s", i, constantKey(co))
		if i < len(c.constants)-1 {
			buf.WriteByte('\n')
		}
	}

	return buf.Bytes()
}

func constantKey(value any) string {
	return fmt.Sprintf("%T:%#v", value, value)
}
