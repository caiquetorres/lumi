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
		// case emitter.FnDecl:
		// 	nameIdx, err := c.readUint32()
		// 	if err != nil {
		// 		return fmt.Errorf("invalid function declaration name index at pc=%d: %w", c.pc, err)
		// 	}

		// 	entryPoint, err := c.readUint32()
		// 	if err != nil {
		// 		return fmt.Errorf("invalid function declaration entry point at pc=%d: %w", c.pc, err)
		// 	}

		// 	if err := m.registerFunction(nameIdx, entryPoint); err != nil {
		// 		return err
		// 	}

		// case emitter.LoadConst, emitter.DefineSymbol, emitter.SetSymbol:
		// 	if _, err := c.readUint32(); err != nil {
		// 		return fmt.Errorf("invalid uint32 operand for opcode %d at pc=%d: %w", opcode, c.pc, err)
		// 	}

		case emitter.Call:
			// Skip the call arity operand (1 byte)

			if _, err := c.readUint8(); err != nil {
				return fmt.Errorf("invalid call arity operand at pc=%d: %w", c.pc, err)
			}

		case emitter.Pop,
			emitter.Return, emitter.Add, emitter.Sub, emitter.Mul, emitter.Div, emitter.Eq, emitter.Not, emitter.Less, emitter.LessEq:
			// No operands to consume.

		case emitter.JumpTo:
			if _, err := c.readUint32(); err != nil {
				return fmt.Errorf("invalid jump offset operand at pc=%d: %w", c.pc, err)
			}

		case emitter.JumpIfFalse:
			if _, err := c.readUint32(); err != nil {
				return fmt.Errorf("invalid jump offset operand at pc=%d: %w", c.pc, err)
			}

		default:
			return fmt.Errorf("unknown opcode %d at pc=%d", opcode, c.pc-1)
		}
	}

	m.registerNativeFunctions()

	return nil
}

func (m *vm) registerNativeFunctions() {
	m.globals.define("printf", nativeFn{
		fn: func(args ...any) (any, error) {
			fmt.Printf(args[0].(string), args[1:]...)
			return nil, nil
		},
	})

	m.globals.define("println", nativeFn{
		fn: func(args ...any) (any, error) {
			fmt.Println(args...)
			return nil, nil
		},
	})

	m.globals.define("sprintf", nativeFn{
		fn: func(args ...any) (any, error) {
			return fmt.Sprintf(args[0].(string), args[1:]...), nil
		},
	})

	m.globals.define("len", nativeFn{
		fn: func(args ...any) (any, error) {
			str, ok := args[0].(string)
			if !ok {
				return nil, fmt.Errorf("len expects a string argument")
			}

			return len(str), nil
		},
	})
}

func (m *vm) registerFunction(nameIdx uint32, entryPoint uint32) error {
	fnName, err := m.pool.GetConstantAsString(nameIdx)
	if err != nil {
		return err
	}

	m.globals.define(fnName, fn{
		entry: entryPoint,
	})

	return nil
}
