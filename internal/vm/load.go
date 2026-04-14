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

			entryPoint, nextPC, err := m.readUint32At(pc)
			if err != nil {
				return fmt.Errorf("invalid function declaration entry point at pc=%d: %w", pc, err)
			}
			pc = nextPC

			if err := m.registerFunction(nameIdx, entryPoint); err != nil {
				return err
			}

		case emitter.LoadConst, emitter.GetSymbol:
			_, nextPC, err := m.readUint32At(pc)
			if err != nil {
				return fmt.Errorf("invalid uint32 operand for opcode %d at pc=%d: %w", opcode, pc, err)
			}
			pc = nextPC

		case emitter.End, emitter.BeginScope, emitter.EndScope,
			emitter.Call, emitter.Pop:
			// No operands to consume.

		default:
			return fmt.Errorf("unknown opcode %d at pc=%d", opcode, pc-1)
		}
	}

	return nil
}

func (m *vm) registerFunction(nameIdx uint32, entryPoint uint32) error {
	name, exists := m.c.getConstant(nameIdx)
	if !exists {
		return fmt.Errorf("constant with index %d not found", nameIdx)
	}

	fnName, ok := name.(string)
	if !ok {
		return fmt.Errorf("expected string constant for function name, got %T", name)
	}

	m.symbolTable.define(fnName, fn{
		name:  fnName,
		entry: entryPoint,
	})

	return nil
}
