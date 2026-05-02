package vm

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"

	"github.com/caiquetorres/lumi/internal/constpool"
)

type vm struct {
	src []byte

	pool         *constpool.ConstantPool
	frames       *frames
	operandStack *operandStack
	heap         *heap
}

func Exec(src io.Reader) error {
	r := bufio.NewReader(src)

	if !isLumiFile(r) {
		return errors.New("not a lumi file")
	}

	constants, err := getConstants(r)
	if err != nil {
		return err
	}

	pool, err := constpool.ParseConstantPool(constants)
	if err != nil {
		return err
	}

	entryPoint, hasEntryPoint, err := getEntryPoint(r)
	if err != nil {
		return err
	}

	instructions, err := getInstructions(r)
	if err != nil {
		return err
	}

	machine := &vm{
		src:          instructions,
		pool:         pool,
		frames:       newFrames(),
		operandStack: newOperandStack(1024),
		heap:         newHeap(4 * 1024 * 1024), // 4 MB heap
	}

	if err := machine.load(0); err != nil {
		return err
	}

	machine.frames.reset()

	if hasEntryPoint {
		return machine.run(entryPoint)
	}

	return nil
}

const lumiMagic = "LUMI"

func isLumiFile(fp *bufio.Reader) bool {
	var magic [4]byte

	n, err := fp.Read(magic[:])
	if err != nil || n != 4 {
		return false
	}

	return string(magic[:]) == lumiMagic
}

func getConstantPoolLen(fp *bufio.Reader) (int, error) {
	var sizeBuf [4]byte

	n, err := fp.Read(sizeBuf[:])
	if err != nil || n != 4 {
		return 0, err
	}

	return int(binary.BigEndian.Uint32(sizeBuf[:])), nil
}

func getConstants(fp *bufio.Reader) ([]byte, error) {
	size, err := getConstantPoolLen(fp)
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

func getEntryPoint(fp *bufio.Reader) (uint32, bool, error) {
	b, err := fp.ReadByte()
	if err != nil {
		return 0, false, err
	}

	if b != 1 {
		return 0, false, nil
	}

	var entryPointBuf [4]byte
	if _, err := fp.Read(entryPointBuf[:]); err != nil {
		return 0, false, err
	}

	entryPoint := binary.BigEndian.Uint32(entryPointBuf[:])
	return entryPoint, true, nil
}

func getInstructions(fp *bufio.Reader) ([]byte, error) {
	instructions, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}
