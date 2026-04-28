package vm

import (
	"fmt"
)

type vm struct {
	c *constantPool

	globals     *symbolTable
	symbolTable *symbolTable

	stack  []any
	frames frames
	src    []byte
}

type fn struct {
	entry uint32
}

type nativeFn struct {
	fn func(args ...any) (any, error)
}

func (m *vm) readConstant() (any, error) {
	idx, err := m.frames.current().readUint32()
	if err != nil {
		return 0, err
	}

	constant, exists := m.c.getConstant(idx)
	if !exists {
		return nil, fmt.Errorf("constant with index %d not found", idx)
	}

	return constant, nil
}

func (m *vm) pushObject(obj any) {
	m.stack = append(m.stack, obj)
}

func (m *vm) popObject() (any, error) {
	if len(m.stack) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	obj := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]

	return obj, nil
}
