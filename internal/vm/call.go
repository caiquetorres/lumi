package vm

import (
	"fmt"
	"slices"
)

func (m *vm) skipCall(pc *uint32) error {
	// Skip the call arity operand (1 byte)
	_, nextPC, err := m.readUint8At(*pc)
	if err != nil {
		return err
	}

	*pc = nextPC

	return nil
}

func (m *vm) execCall() error {
	obj, err := m.popObject()
	if err != nil {
		return err
	}

	arity, err := m.readUint8()
	if err != nil {
		return err
	}

	switch fnObj := obj.(type) {
	case fn:
		return m.callFn(&fnObj, arity)
	case nativeFn:
		return m.callNativeFn(&fnObj, arity)
	}

	return nil
}

func (m *vm) callFn(fnObj *fn, arity uint8) error {
	params := make(map[string]any, len(fnObj.params))

	for i := int(arity) - 1; i >= 0; i-- {
		if i >= len(fnObj.params) {
			continue
		}

		paramNameConst, exists := m.c.getConstant(fnObj.params[i])
		if !exists {
			return fmt.Errorf("constant with index %d not found", fnObj.params[i])
		}

		paramName, ok := paramNameConst.(string)
		if !ok {
			return fmt.Errorf("expected string constant for parameter name, got %T", paramNameConst)
		}

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

func (m *vm) callNativeFn(fnObj *nativeFn, arity uint8) error {
	args := make([]any, 0, arity)

	for i := int(arity) - 1; i >= 0; i-- {
		arg, err := m.popObject()
		if err != nil {
			return err
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
