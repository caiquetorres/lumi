package vm

import "log"

type operandType byte

const (
	operandInt operandType = iota + 1
	operandBool
	operandString
	operandFn
	operandNativeFn
)

type operand struct {
	ty operandType

	intValue  int64
	boolValue bool
	fnValue   uint32
	strValue  uint64
}

func intOperand(value int64) operand {
	return operand{
		ty:       operandInt,
		intValue: value,
	}
}

func fnOperand(value uint32) operand {
	return operand{
		ty:      operandFn,
		fnValue: value,
	}
}

func nativeFnOperand(value uint32) operand {
	return operand{
		ty:      operandNativeFn,
		fnValue: value,
	}
}

func boolOperand(value bool) operand {
	return operand{
		ty:        operandBool,
		boolValue: value,
	}
}

func stringOperand(addr uint64) operand {
	return operand{
		ty:       operandString,
		strValue: addr,
	}
}

type operandStack struct {
	data []operand
}

func newOperandStack(size int) *operandStack {
	return &operandStack{
		data: make([]operand, 0, size),
	}
}

func (s *operandStack) push(v operand) {
	s.data = append(s.data, v)
}

func (s *operandStack) pop() operand {
	top := len(s.data) - 1
	if top < 0 {
		log.Panic("operand stack underflow: cannot pop from an empty stack")
	}

	v := s.data[top]
	s.data = s.data[:top]
	return v
}
