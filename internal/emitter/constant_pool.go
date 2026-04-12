package emitter

import (
	"bytes"
	"fmt"
)

type constantPool struct {
	constants  []any
	indexByKey map[string]int
}

func newConstantPool() *constantPool {
	return &constantPool{
		constants:  []any{},
		indexByKey: map[string]int{},
	}
}

func (c *constantPool) internConstant(value any) int {
	if c.indexByKey == nil {
		c.indexByKey = map[string]int{}
	}

	key := constantKey(value)
	if idx, ok := c.indexByKey[key]; ok {
		return idx
	}

	c.constants = append(c.constants, value)
	idx := len(c.constants) - 1
	c.indexByKey[key] = idx

	return idx
}

func (c *constantPool) setContant(index int, value any) {
	if index < len(c.constants) {
		c.constants[index] = value
	}
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
