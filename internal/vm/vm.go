package vm

import (
	"encoding/binary"
	"fmt"
	"io"
)

type vm struct {
	c           *constantPool
	symbolTable *symbolTable

	stack  []any
	frames frames
	src    []byte
}

type fn struct {
	name       string
	entry      uint32
	paramNames []string
}

type nativeFn struct {
	name string
	fn   func(args ...any) (any, error)
}

func (m *vm) nextInstruction() (byte, error) {
	if int(m.frames.current()) >= len(m.src) {
		return 0, io.EOF
	}

	b := m.src[m.frames.current()]
	m.frames.incCurrentPtr(1)

	return b, nil
}

func (m *vm) readUint32() (uint32, error) {
	val, _, err := m.readUint32At(m.frames.current())
	if err != nil {
		return 0, err
	}

	m.frames.incCurrentPtr(4)

	return val, nil
}

func (m *vm) readUint32At(pc uint32) (uint32, uint32, error) {
	const uint32Size = 4

	if pc+uint32Size > uint32(len(m.src)) {
		return 0, pc, io.EOF
	}

	value := binary.BigEndian.Uint32(m.src[pc : pc+uint32Size])
	return value, pc + uint32Size, nil
}

func (m *vm) readConstant() (any, error) {
	idx, err := m.readUint32()
	if err != nil {
		return nil, err
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
