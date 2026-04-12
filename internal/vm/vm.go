package vm

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/caiquetorres/lumi/internal/emitter"
)

type vm struct {
	stack  []any
	c      *constantPool
	frames []frame
	src    []byte

	symbolTable map[string]any
}

type frame struct {
	ptr int
}

type fn struct {
	name  string
	entry int
}

func (m *vm) load() error {
	if m.symbolTable == nil {
		m.symbolTable = make(map[string]any)
	}

	scanPtr := 0

	for scanPtr < len(m.src) {
		opcode := m.src[scanPtr]
		scanPtr++

		switch opcode {
		case emitter.DeclFun:
			// Read function name constant index
			if scanPtr+4 > len(m.src) {
				return io.EOF
			}

			nameIdx := int(binary.BigEndian.Uint32(m.src[scanPtr : scanPtr+4]))
			scanPtr += 4

			// Read function entry point address
			if scanPtr+4 > len(m.src) {
				return io.EOF
			}

			entryPoint := int(binary.BigEndian.Uint32(m.src[scanPtr : scanPtr+4]))
			scanPtr += 4

			// Get function name from constant pool
			name, exists := m.c.getConstant(nameIdx)
			if !exists {
				return fmt.Errorf("constant with index %d not found", nameIdx)
			}

			fnName, ok := name.(string)
			if !ok {
				return fmt.Errorf("expected string constant for function name, got %T", name)
			}

			// Store in symbol table
			m.symbolTable[fnName] = fn{
				name:  fnName,
				entry: entryPoint,
			}
		}
	}

	return nil
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

		case emitter.GetSymbol:
			constant, err := m.readConstant()
			if err != nil {
				return err
			}

			name, ok := constant.(string)
			if !ok {
				return fmt.Errorf("expected string constant for symbol name, got %T", constant)
			}

			value, exists := m.symbolTable[name]
			if !exists {
				return fmt.Errorf("symbol %q not found", name)
			}

			m.stack = append(m.stack, value)

		case emitter.Call:
			obj := m.popObject()
			fnObj, ok := obj.(fn)
			if !ok {
				return fmt.Errorf("expected function object on stack, got %T", obj)
			}

			m.frames = append(m.frames, frame{ptr: fnObj.entry})

		case emitter.Pop:
			obj := m.popObject()
			fmt.Println(obj)

		case emitter.End:
			m.frames = m.frames[:len(m.frames)-1]

			if len(m.frames) == 0 {
				return nil
			}

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
