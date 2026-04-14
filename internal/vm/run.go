package vm

import (
	"fmt"
	"slices"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) run() error {
	m.symbolTable.define("printf", nativeFn{
		name: "printf",
		fn: func(args ...any) (any, error) {
			fmt.Printf(args[0].(string), args[1:]...)
			return nil, nil
		},
	})

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

			value, exists := m.symbolTable.lookup(name)
			if !exists {
				return fmt.Errorf("symbol %q not found", name)
			}

			m.pushObject(value)

		case emitter.Call:
			obj := m.popObject()

			switch fnObj := obj.(type) {
			case fn:
				params := make(map[string]any, len(fnObj.paramNames))

				for i := len(fnObj.paramNames) - 1; i >= 0; i-- {
					paramName := fnObj.paramNames[i]
					paramValue := m.popObject()

					params[paramName] = paramValue
				}

				m.frames.push(fnObj.entry)
				m.symbolTable = newSymbolTable(m.symbolTable)
				for name, value := range params {
					m.symbolTable.define(name, value)
				}
			case nativeFn:
				args := make([]any, 0)

				for {
					arg := m.popObject()
					if arg == nil {
						break
					}

					args = append(args, arg)
				}

				slices.Reverse(args)

				result, err := fnObj.fn(args...)
				if err != nil {
					return fmt.Errorf("error calling native function %q: %w", fnObj.name, err)
				}

				m.pushObject(result)
			}

		case emitter.Pop:
			_ = m.popObject()

		case emitter.End:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}
		}
	}
}
