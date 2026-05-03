package vm

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/caiquetorres/lumi/internal/constpool"
)

type nativeFn func(m *vm, args ...operand) (operand, error)

type vm struct {
	src []byte

	pool          *constpool.ConstantPool
	frames        *frames
	operandStack  *operandStack
	fnTable       []uint32
	nativeFnTable map[string]nativeFn
	heap          *heap
	locals        [1024]byte
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

	fnTable, err := getFunctionTable(r)
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

	nativeFnTable := map[string]nativeFn{
		"println": func(m *vm, args ...operand) (operand, error) {
			values := make([]any, len(args))
			for i, arg := range args {
				switch arg.ty {
				case operandInt:
					values[i] = arg.intValue
				case operandBool:
					values[i] = arg.boolValue
				case operandString:
					strAddr := decodeString(arg.strValue)
					strObj, err := m.heap.readObject(strAddr)
					if err != nil {
						return operand{}, fmt.Errorf("failed to read string from heap: %w", err)
					}
					values[i] = string(strObj.data)
				default:
					return operand{}, fmt.Errorf("unsupported operand type for println: %v", arg.ty)
				}
			}
			fmt.Println(values...)
			return operand{}, nil
		},
		"printf": func(m *vm, args ...operand) (operand, error) {
			if len(args) == 0 {
				return operand{}, errors.New("printf requires at least a format string")
			}

			formatArg := args[0]
			if formatArg.ty != operandString {
				return operand{}, errors.New("first argument to printf must be a string")
			}

			formatAddr := decodeString(formatArg.strValue)
			formatObj, err := m.heap.readObject(formatAddr)
			if err != nil {
				return operand{}, fmt.Errorf("failed to read format string from heap: %w", err)
			}

			formatStr := string(formatObj.data)

			values := make([]any, len(args)-1)
			for i, arg := range args[1:] {
				switch arg.ty {
				case operandInt:
					values[i] = arg.intValue
				case operandBool:
					values[i] = arg.boolValue
				case operandString:
					strAddr := decodeString(arg.strValue)
					strObj, err := m.heap.readObject(strAddr)
					if err != nil {
						return operand{}, fmt.Errorf("failed to read string from heap: %w", err)
					}
					values[i] = string(strObj.data)
				default:
					return operand{}, fmt.Errorf("unsupported operand type for printf: %v", arg.ty)
				}
			}

			fmt.Printf(formatStr, values...)
			return operand{}, nil
		},
	}

	machine := &vm{
		src:           instructions,
		pool:          pool,
		frames:        newFrames(),
		operandStack:  newOperandStack(1024),
		fnTable:       fnTable,
		nativeFnTable: nativeFnTable,
		heap:          newHeap(4 * 1024 * 1024), // 4 MB heap
	}

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

func getFunctionTable(fp *bufio.Reader) ([]uint32, error) {
	var sizeBuf [4]byte

	n, err := fp.Read(sizeBuf[:])
	if err != nil || n != 4 {
		return nil, err
	}

	size := binary.BigEndian.Uint32(sizeBuf[:])
	fnTable := make([]uint32, size)

	for i := uint32(0); i < size; i++ {
		var offsetBuf [4]byte
		if _, err := fp.Read(offsetBuf[:]); err != nil {
			return nil, err
		}

		fnTable[i] = binary.BigEndian.Uint32(offsetBuf[:])
	}

	return fnTable, nil
}

func getInstructions(fp *bufio.Reader) ([]byte, error) {
	instructions, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}
