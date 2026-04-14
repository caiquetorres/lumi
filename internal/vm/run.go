package vm

import (
	"fmt"
	"slices"

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

		case emitter.End:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}
		}
	}
}

func (m *vm) execCall() error {
	obj, err := m.popObject()
	if err != nil {
		return err
	}

	switch fnObj := obj.(type) {
	case fn:
		return m.callFn(&fnObj)
	case nativeFn:
		return m.callNativeFn(fnObj)
	}

	return nil
}

func (m *vm) callFn(fnObj *fn) error {
	params := make(map[string]any, len(fnObj.paramNames))

	for i := len(fnObj.paramNames) - 1; i >= 0; i-- {
		paramName := fnObj.paramNames[i]
		paramValue, err := m.popObject()
		if err != nil {
			return err
		}

		params[paramName] = paramValue
	}

	m.frames.push(fnObj.entry)
	m.symbolTable = newSymbolTable(m.symbolTable)

	for name, value := range params {
		m.symbolTable.define(name, value)
	}

	return nil
}

func (m *vm) callNativeFn(fnObj nativeFn) error {
	args := make([]any, 0)

	for {
		arg, _ := m.popObject()
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

	return nil
}
