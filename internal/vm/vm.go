package vm

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/caiquetorres/lumi/internal/emitter"
)

type vm struct {
	ptr int // instruction pointer

	stack      []any
	constStack []int

	c   *constantPool
	src []byte
}

func Execute(src io.ReadSeeker) error {
	if !isLumiFile(src) {
		return nil
	}

	constants, err := getConstants(src)
	if err != nil {
		return err
	}

	c, err := parseConstantPool(constants)
	if err != nil {
		return err
	}

	entryPoint, hasEntryPoint, err := getEntryPoint(src)
	if err != nil {
		return err
	}

	if !hasEntryPoint {
		return fmt.Errorf("no entry point found")
	}

	instructions, err := getInstructions(src)
	if err != nil {
		return err
	}

	machine := &vm{
		ptr: entryPoint,
		c:   c,
		src: instructions,
	}

	return machine.run()
}

func (m *vm) run() error {
	for {
		i, err := m.nextInstruction()
		if err != nil {
			return err
		}

		switch i {
		// case emitter.LoadConst:
		// 	idx, err := m.readInt()
		// 	if err != nil {
		// 		return err
		// 	}

		// 	m.constStack = append(m.constStack, idx)

		// case emitter.DeclFun:
		// 	// load functions's name
		// 	fnNameIdx := m.constStack[len(m.constStack)-1]
		// 	m.constStack = m.constStack[:len(m.constStack)-1]

		// 	fnName := m.c.constants[fnNameIdx].(string)

		// 	fnPtr, err := m.readInt()
		// 	if err != nil {
		// 		return err
		// 	}

		// 	fmt.Printf("Declared function %s at %d\n", fnName, fnPtr)

		case emitter.End:
			return nil
		}
	}
}

func (m *vm) nextInstruction() (byte, error) {
	if m.ptr >= len(m.src) {
		return 0, io.EOF
	}

	b := m.src[m.ptr]
	m.ptr++
	return b, nil
}

func (m *vm) readInt() (int, error) {
	if m.ptr+4 > len(m.src) {
		return 0, io.EOF
	}

	buf := m.src[m.ptr : m.ptr+4]
	m.ptr += 4
	return int(binary.BigEndian.Uint32(buf)), nil
}

func isLumiFile(fp io.ReadSeeker) bool {
	magic := make([]byte, 4)

	n, err := fp.Read(magic)
	if err != nil || n != 4 {
		return false
	}

	return string(magic) == "LUMI"
}

func getConstantsSize(fp io.ReadSeeker) (int, error) {
	sizeBuf := make([]byte, 4)

	n, err := fp.Read(sizeBuf)
	if err != nil || n != 4 {
		return 0, err
	}

	return int(binary.BigEndian.Uint32(sizeBuf)), nil
}

func getConstants(fp io.ReadSeeker) ([]byte, error) {
	size, err := getConstantsSize(fp)
	if err != nil {
		return nil, err
	}

	constants := make([]byte, size)

	n, err := fp.Read(constants)
	if err != nil || n != size {
		return nil, err
	}

	return constants, nil
}

func getEntryPoint(fp io.ReadSeeker) (int, bool, error) {
	entryPointFlag := make([]byte, 1)

	n, err := fp.Read(entryPointFlag)
	if err != nil || n != 1 {
		return 0, false, err
	}

	if entryPointFlag[0] == 0 {
		return 0, false, nil
	}

	entryPointBuf := make([]byte, 4)

	n, err = fp.Read(entryPointBuf)
	if err != nil || n != 4 {
		return 0, false, err
	}

	return int(binary.BigEndian.Uint32(entryPointBuf)), true, nil
}

func getInstructions(fp io.ReadSeeker) ([]byte, error) {
	instructions, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}
