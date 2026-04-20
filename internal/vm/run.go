package vm

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) run() error {
	mainFnObj, hasMain := m.globals.lookup("main")
	if !hasMain {
		return fmt.Errorf("no main function found")
	}

	mainFn := mainFnObj.(fn)
	if err := m.callFn(&mainFn, 0); err != nil {
		return err
	}

	for {
		opcode, err := m.frames.current().readUint8()
		if err != nil {
			return err
		}

		switch opcode {
		case emitter.LoadConst:
			constant, err := m.readConstant()
			if err != nil {
				return err
			}

			// REVIEW: How does it know that the constant has a value that can be pushed onto the stack?

			m.pushObject(constant)

		case emitter.VarDecl:
			constant, err := m.readConstant()
			if err != nil {
				return err
			}

			name, ok := constant.(string)
			if !ok {
				return fmt.Errorf("expected string constant for symbol name, got %T", constant)
			}

			value, err := m.popObject()
			if err != nil {
				return err
			}

			m.symbolTable.define(name, value)

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
			if err := m.execCall(); err != nil {
				return err
			}

		case emitter.Pop:
			_, _ = m.popObject()

		case emitter.Jump:
			offset, err := m.frames.current().readUint32()
			if err != nil {
				return fmt.Errorf("invalid jump offset operand: %w", err)
			}

			m.frames.current().moveTo(offset)

		case emitter.Return:

		case emitter.BeginScope:
			m.symbolTable = newSymbolTable(m.symbolTable)

		case emitter.EndScope:
			if m.symbolTable.parent != nil {
				m.symbolTable = m.symbolTable.parent
			}

		case emitter.End:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}
		}
	}
}
