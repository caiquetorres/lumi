package vm

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/emitter"
)

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
