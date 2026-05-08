package vm

import "fmt"

func (m *vm) call() error {
	val := m.operandStack.pop()

	if val.ty != operandFn && val.ty != operandNativeFn {
		return fmt.Errorf("expected function operand for Call, got %v", val.ty)
	}

	arity, err := m.frames.current().readUint8(m.src)
	if err != nil {
		return fmt.Errorf("failed to read arity for Call: %w", err)
	}

	switch val.ty {
	case operandFn:
		return m.callFn(&val)
	case operandNativeFn:
		return m.callNativeFn(&val, arity)
	}

	// unreachable
	return nil
}

func (m *vm) callFn(op *operand) error {
	tmp := m.frames.current().tmp
	m.frames.push(op.fnValue, tmp)

	return nil
}

func (m *vm) callNativeFn(op *operand, arity uint8) error {
	fnNameIndex := op.fnValue
	fnNameConst, exists := m.pool.GetConstant(fnNameIndex)
	if !exists {
		return fmt.Errorf("failed to get function name constant")
	}

	operands := make([]operand, arity)
	for i := range operands {
		operands[i] = m.operandStack.pop()
	}

	fn := m.nativeFnTable[fnNameConst.(string)]

	res, err := fn(m, operands...)
	if err != nil {
		return fmt.Errorf("error calling native function: %w", err)
	}

	m.operandStack.push(res)

	return nil
}
