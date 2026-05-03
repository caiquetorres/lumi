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

	if val.ty == operandFn {
		tmp := m.frames.current().tmp
		m.frames.push(val.fnValue, tmp, m.src)

		return nil
	}

	fnNameIndex := val.fnValue
	fnNameConst, exists := m.pool.GetConstant(fnNameIndex)
	if !exists {
		return fmt.Errorf("failed to get function name constant")
	}

	operands := make([]operand, arity)
	for i := int(arity) - 1; i >= 0; i-- {
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
