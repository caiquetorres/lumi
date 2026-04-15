package vm

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) load() error {
	if m.symbolTable == nil {
		m.symbolTable = newSymbolTable(nil)
	}

	pc := uint32(0)

	for pc < uint32(len(m.src)) {
		opcode := m.src[pc]
		pc++

		switch opcode {
		case emitter.DeclFun:
			nameIdx, nextPC, err := m.readUint32At(pc)
			if err != nil {
				return fmt.Errorf("invalid function declaration name index at pc=%d: %w", pc, err)
			}
			pc = nextPC

			paramCount, nextPC, err := m.readUint8At(pc)
			if err != nil {
				return fmt.Errorf("invalid function declaration parameter count at pc=%d: %w", pc, err)
			}
			pc = nextPC

			var params []uint32
			for range paramCount {
				paramIdx, nextPC, err := m.readUint32At(pc)
				if err != nil {
					return fmt.Errorf("invalid function declaration parameter index at pc=%d: %w", pc, err)
				}
				params = append(params, paramIdx)
				pc = nextPC
			}

			entryPoint, nextPC, err := m.readUint32At(pc)
			if err != nil {
				return fmt.Errorf("invalid function declaration entry point at pc=%d: %w", pc, err)
			}
			pc = nextPC

			err = m.registerFunction(nameIdx, params, entryPoint)
			if err != nil {
				return err
			}
			pc = nextPC

		case emitter.LoadConst, emitter.GetSymbol:
			_, nextPC, err := m.readUint32At(pc)
			if err != nil {
				return fmt.Errorf("invalid uint32 operand for opcode %d at pc=%d: %w", opcode, pc, err)
			}
			pc = nextPC

		case emitter.Call:
			if err := m.skipCall(&pc); err != nil {
				return fmt.Errorf("failed to skip operands for opcode %d at pc=%d: %w", opcode, pc, err)
			}

		case emitter.End, emitter.BeginScope, emitter.EndScope, emitter.Pop:
			// No operands to consume.

		default:
			return fmt.Errorf("unknown opcode %d at pc=%d", opcode, pc-1)
		}
	}

	m.symbolTable.define("printf", nativeFn{
		fn: func(args ...any) (any, error) {
			fmt.Printf(args[0].(string), args[1:]...)
			return nil, nil
		},
	})

	return nil
}

func (m *vm) registerFunction(nameIdx uint32, params []uint32, entryPoint uint32) error {
	fnNameObj, exists := m.c.getConstant(nameIdx)
	if !exists {
		return fmt.Errorf("constant with index %d not found", nameIdx)
	}

	fnName, ok := fnNameObj.(string)
	if !ok {
		return fmt.Errorf("expected string constant for function name, got %T", fnNameObj)
	}

	m.symbolTable.define(fnName, fn{
		entry:  entryPoint,
		params: params,
	})

	return nil
}
