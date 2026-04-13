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
	frames frames
	src    []byte

	symbolTable map[string]any
}

type fn struct {
	name  string
	entry uint32
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

			nameIdx := binary.BigEndian.Uint32(m.src[scanPtr : scanPtr+4])
			scanPtr += 4

			// Read function entry point address
			if scanPtr+4 > len(m.src) {
				return io.EOF
			}

			entryPoint := binary.BigEndian.Uint32(m.src[scanPtr : scanPtr+4])
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

			// REVIEW: How does it know that the constant has a value that can be pushed onto the stack?

			m.pushObject(constant)

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

			m.pushObject(value)

		case emitter.Call:
			obj := m.popObject()
			fnObj, ok := obj.(fn)
			if !ok {
				return fmt.Errorf("expected function object on stack, got %T", obj)
			}

			m.frames.push(fnObj.entry)

		case emitter.Pop:
			obj := m.popObject()
			if obj != nil {
				fmt.Println(obj)
			}

		case emitter.End:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}
		}
	}
}

func (m *vm) nextInstruction() (byte, error) {
	if m.frames.current().ptr >= uint32(len(m.src)) {
		return 0, io.EOF
	}

	b := m.src[m.frames.current().ptr]
	m.frames.incrementCurrentPtr(1)

	return b, nil
}

func (m *vm) readUint32() (uint32, error) {
	if m.frames.current().ptr+4 > uint32(len(m.src)) {
		return 0, io.EOF
	}

	ptr := m.frames.current().ptr

	buf := m.src[ptr : ptr+4]
	m.frames.incrementCurrentPtr(4)

	return binary.BigEndian.Uint32(buf), nil
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

func (m *vm) popObject() any {
	if len(m.stack) == 0 {
		return nil
	}

	obj := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]

	return obj
}
