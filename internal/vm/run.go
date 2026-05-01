package vm

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/emitter"
)

// *(*interface {})(0x14000128138)

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

		case emitter.DefineSymbol:
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

			if val, ok := value.(int); ok {
				value = val
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

			if val, ok := value.(int); ok {
				value = val
			}

			// copy the value to the stack so that it can be used by subsequent instructions without modifying the symbol table entry

			m.pushObject(value)

		case emitter.Call:
			if err := m.execCall(); err != nil {
				return err
			}

		case emitter.Pop:
			_, _ = m.popObject()

		case emitter.JumpTo:
			offset, err := m.frames.current().readUint32()
			if err != nil {
				return fmt.Errorf("invalid jump offset operand: %w", err)
			}

			m.frames.current().moveTo(offset)

		case emitter.JumpIfFalse:
			offset, err := m.frames.current().readUint32()
			if err != nil {
				return fmt.Errorf("invalid jump offset operand: %w", err)
			}

			condition, err := m.popObject()
			if err != nil {
				return err
			}

			switch v := condition.(type) {
			case bool:
				condition = v
			case string:
				condition = v != ""
			default:
				return fmt.Errorf("expected boolean condition for JumpIfFalse, got %T", condition)
			}

			isFalse := condition == false
			if isFalse {
				m.frames.current().moveTo(offset)
			}

		case emitter.Return:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}

		case emitter.Add:
			right, err := m.popObject()
			if err != nil {
				return err
			}

			left, err := m.popObject()
			if err != nil {
				return err
			}

			// REVIEW: This is a very naive implementation. It only supports adding integers. We will need to add type checking and support for other types (e.g., strings) in the future.

			leftInt, rightInt := left.(int), right.(int)
			m.pushObject(leftInt + rightInt)

		case emitter.Sub:
			right, err := m.popObject()
			if err != nil {
				return err
			}

			left, err := m.popObject()
			if err != nil {
				return err
			}

			leftInt, rightInt := left.(int), right.(int)
			m.pushObject(leftInt - rightInt)

		case emitter.Mul:
			right, err := m.popObject()
			if err != nil {
				return err
			}

			left, err := m.popObject()
			if err != nil {
				return err
			}

			leftInt, rightInt := left.(int), right.(int)
			m.pushObject(leftInt * rightInt)

		case emitter.Div:
			right, err := m.popObject()
			if err != nil {
				return err
			}

			left, err := m.popObject()
			if err != nil {
				return err
			}

			leftInt, rightInt := left.(int), right.(int)
			m.pushObject(leftInt / rightInt)

		case emitter.Eq:
			right, err := m.popObject()
			if err != nil {
				return err
			}

			left, err := m.popObject()
			if err != nil {
				return err
			}

			m.pushObject(left == right)

		case emitter.Not:
			value, err := m.popObject()
			if err != nil {
				return err
			}

			boolValue, ok := value.(bool)
			if !ok {
				return fmt.Errorf("expected boolean value for Not operation, got %T", value)
			}

			m.pushObject(!boolValue)

		case emitter.Less:
			right, err := m.popObject()
			if err != nil {
				return err
			}

			left, err := m.popObject()
			if err != nil {
				return err
			}

			leftInt, rightInt := left.(int), right.(int)
			m.pushObject(leftInt < rightInt)

		case emitter.LessEq:
			right, err := m.popObject()
			if err != nil {
				return err
			}

			left, err := m.popObject()
			if err != nil {
				return err
			}

			leftInt, rightInt := left.(int), right.(int)
			m.pushObject(leftInt <= rightInt)

		case emitter.SetSymbol:
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

			if val, ok := value.(int); ok {
				value = val
			}

			if !m.symbolTable.set(name, value) {
				return fmt.Errorf("symbol %q not found for assignment", name)
			}

			m.pushObject(value)

		case emitter.BeginScope:
			m.symbolTable = newSymbolTable(m.symbolTable)

		case emitter.EndScope:
			if m.symbolTable.parent != nil {
				m.symbolTable = m.symbolTable.parent
			}
		}
	}
}
