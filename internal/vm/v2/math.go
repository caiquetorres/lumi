package vm

import "fmt"

func (m *vm) add() error {
	right := m.operandStack.pop()
	left := m.operandStack.pop()

	if right.ty != operandInt || left.ty != operandInt {
		return fmt.Errorf("operands for Add must be integers")
	}

	result := left.intValue + right.intValue
	m.operandStack.push(intOperand(result))

	return nil
}

func (m *vm) sub() error {
	right := m.operandStack.pop()
	left := m.operandStack.pop()

	if right.ty != operandInt || left.ty != operandInt {
		return fmt.Errorf("operands for Sub must be integers")
	}

	result := left.intValue - right.intValue
	m.operandStack.push(intOperand(result))

	return nil
}

func (m *vm) mul() error {
	right := m.operandStack.pop()
	left := m.operandStack.pop()

	if right.ty != operandInt || left.ty != operandInt {
		return fmt.Errorf("operands for Mul must be integers")
	}

	result := left.intValue * right.intValue
	m.operandStack.push(intOperand(result))

	return nil
}

func (m *vm) div() error {
	right := m.operandStack.pop()
	left := m.operandStack.pop()

	if right.ty != operandInt || left.ty != operandInt {
		return fmt.Errorf("operands for Div must be integers")
	}

	if right.intValue == 0 {
		return fmt.Errorf("division by zero")
	}

	result := left.intValue / right.intValue
	m.operandStack.push(intOperand(result))

	return nil
}

func (m *vm) eq() error {
	right := m.operandStack.pop()
	left := m.operandStack.pop()

	if right.ty != left.ty {
		return fmt.Errorf("operands for Eq must be of the same type")
	}

	var result bool
	switch right.ty {
	case operandInt:
		result = left.intValue == right.intValue

	case operandBool:
		result = left.boolValue == right.boolValue

	case operandString:
		addr := decodeString(left.stringValue)
		leftStr, err := m.heap.readObject(addr)

		if err != nil {
			return fmt.Errorf("failed to read string object: %w", err)
		}

		addr = decodeString(right.stringValue)
		rightStr, err := m.heap.readObject(addr)

		if err != nil {
			return fmt.Errorf("failed to read string object: %w", err)
		}

		result = string(leftStr.data) == string(rightStr.data)

	case operandFn:
		result = left.fnValue == right.fnValue

	default:
		return fmt.Errorf("unsupported operand type for Eq: %v", right.ty)
	}

	op := boolOperand(result)
	m.operandStack.push(op)

	return nil
}

func (m *vm) less() error {
	right := m.operandStack.pop()
	left := m.operandStack.pop()

	if right.ty != operandInt || left.ty != operandInt {
		return fmt.Errorf("operands for Less must be integers")
	}

	result := left.intValue < right.intValue
	op := boolOperand(result)
	m.operandStack.push(op)

	return nil
}

func (m *vm) lessEq() error {
	right := m.operandStack.pop()
	left := m.operandStack.pop()

	if right.ty != operandInt || left.ty != operandInt {
		return fmt.Errorf("operands for LessEq must be integers")
	}

	result := left.intValue <= right.intValue
	op := boolOperand(result)
	m.operandStack.push(op)

	return nil
}

func (m *vm) not() error {
	value := m.operandStack.pop()

	if value.ty != operandBool {
		return fmt.Errorf("operand for Not must be a boolean")
	}

	val := !value.boolValue
	op := boolOperand(val)
	m.operandStack.push(op)

	return nil
}
