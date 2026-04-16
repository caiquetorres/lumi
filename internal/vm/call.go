package vm

import "fmt"

func (m *vm) execCall() error {
	obj, err := m.popObject()
	if err != nil {
		return err
	}

	arity, err := m.frames.current().readUint8()
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

		paramName, err := m.c.getConstantAsString(fnObj.params[i])
		if err != nil {
			return err
		}

		paramValue, err := m.popObject()
		if err != nil {
			return err
		}

		params[paramName] = paramValue
	}

	m.frames.push(fnObj.entry, m.src)
	m.symbolTable = newSymbolTable(m.globals)

	for name, value := range params {
		m.symbolTable.define(name, value)
	}

	return nil
}

func (m *vm) callNativeFn(fnObj *nativeFn, arity uint8) error {
	args := make([]any, arity)
	for i := int(arity) - 1; i >= 0; i-- {
		arg, err := m.popObject()
		if err != nil {
			return err
		}

		args[i] = arg
	}

	result, err := fnObj.fn(args...)
	if err != nil {
		return fmt.Errorf("error calling native function: %w", err)
	}

	m.pushObject(result)

	return nil
}
