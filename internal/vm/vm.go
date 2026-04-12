package vm

import (
	"encoding/binary"
	"io"

	"github.com/caiquetorres/lumi/internal/emitter"
)

type vm struct {
	stack      []any
	constStack []int

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

func (m *vm) popConstant() any {
	if len(m.constStack) == 0 {
		return nil
	}

	idx := m.constStack[len(m.constStack)-1]
	m.constStack = m.constStack[:len(m.constStack)-1]

	return m.c.constants[idx]
}
