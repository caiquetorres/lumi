package vm

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/caiquetorres/lumi/internal/emitter"
)

type vm struct {
	stack []any

	c      *constantPool
	src    []byte
	frames []frame
}

type frame struct {
	ptr int
}

func (m *vm) run() error {
	for {
		i, err := m.nextInstruction()
		if err != nil {
			return err
		}

		switch i {
		case emitter.LoadConst:
			constant, err := m.readConstant()
			if err != nil {
				return err
			}

			m.stack = append(m.stack, constant)

		case emitter.Pop:
			obj := m.popObject()
			fmt.Println(obj)

		case emitter.End:
			return nil
		}
	}
}

func (m *vm) ptr() int {
	return m.frames[len(m.frames)-1].ptr
}

func (m *vm) setPtr(ptr int) {
	m.frames[len(m.frames)-1].ptr = ptr
}

func (m *vm) nextInstruction() (byte, error) {
	if m.ptr() >= len(m.src) {
		return 0, io.EOF
	}

	b := m.src[m.ptr()]
	m.setPtr(m.ptr() + 1)
	return b, nil
}

func (m *vm) readInt() (int, error) {
	if m.ptr()+4 > len(m.src) {
		return 0, io.EOF
	}

	buf := m.src[m.ptr() : m.ptr()+4]
	m.setPtr(m.ptr() + 4)
	return int(binary.BigEndian.Uint32(buf)), nil
}

func (m *vm) readConstant() (any, error) {
	idx, err := m.readInt()
	if err != nil {
		return nil, err
	}

	c, exists := m.c.getConstant(idx)
	if !exists {
		return nil, fmt.Errorf("constant with index %d not found", idx)
	}
	return c, nil
}

func (m *vm) popObject() any {
	if len(m.stack) == 0 {
		return nil
	}

	obj := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]

	return obj
}
