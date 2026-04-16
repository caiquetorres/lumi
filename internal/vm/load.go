package vm

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) load() error {
	if m.symbolTable == nil {
		m.symbolTable = newSymbolTable(nil)
	}

	c := newCursor(m.src)

	for c.hasMore() {
		opcode, err := c.readUint8()
		if err != nil {
			return fmt.Errorf("failed to read opcode at pc=%d: %w", c.pc, err)
		}

		switch opcode {
		case emitter.FnDecl:
			nameIdx, err := c.readUint32()
			if err != nil {
				return fmt.Errorf("invalid function declaration name index at pc=%d: %w", c.pc, err)
			}

			paramCount, err := c.readUint8()
			if err != nil {
				return fmt.Errorf("invalid function declaration parameter count at pc=%d: %w", c.pc, err)
			}

			var params []uint32
			for range paramCount {
				paramIdx, err := c.readUint32()
				if err != nil {
					return fmt.Errorf("invalid function declaration parameter index at pc=%d: %w", c.pc, err)
				}
				params = append(params, paramIdx)
			}

			entryPoint, err := c.readUint32()
			if err != nil {
				return fmt.Errorf("invalid function declaration entry point at pc=%d: %w", c.pc, err)
			}

			if err := m.registerFunction(nameIdx, params, entryPoint); err != nil {
				return err
			}

		case emitter.LoadConst, emitter.GetSymbol, emitter.VarDecl:
			if _, err := c.readUint32(); err != nil {
				return fmt.Errorf("invalid uint32 operand for opcode %d at pc=%d: %w", opcode, c.pc, err)
			}

		case emitter.Call:
			// Skip the call arity operand (1 byte)

			if _, err := c.readUint8(); err != nil {
				return fmt.Errorf("invalid call arity operand at pc=%d: %w", c.pc, err)
			}

		case emitter.End, emitter.BeginScope, emitter.EndScope, emitter.Pop:
			// No operands to consume.

		default:
			return fmt.Errorf("unknown opcode %d at pc=%d", opcode, c.pc-1)
		}
	}

	m.globals.define("printf", nativeFn{
		fn: func(args ...any) (any, error) {
			fmt.Printf(args[0].(string), args[1:]...)
			return nil, nil
		},
	})

	return nil
}

func (m *vm) registerFunction(nameIdx uint32, params []uint32, entryPoint uint32) error {
	fnName, err := m.c.getConstantAsString(nameIdx)
	if err != nil {
		return err
	}

	m.globals.define(fnName, fn{
		entry:  entryPoint,
		params: params,
	})

	return nil
}
