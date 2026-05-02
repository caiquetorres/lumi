package vm

import "log"

type operandType byte

const (
	operandInt operandType = iota + 1
	operandBool
	operandString
	operandFn
)

type operand struct {
	ty operandType

	intValue  int64
	boolValue bool
	fnValue   uint32
}

func intOperand(value int64) operand {
	return operand{
		ty:       operandInt,
		intValue: value,
	}
}

func boolOperand(value bool) operand {
	return operand{
		ty:        operandBool,
		boolValue: value,
	}
}

type operandStack struct {
	top  int
	data []operand
}

func newOperandStack(size int) *operandStack {
	return &operandStack{
		top:  0,
		data: make([]operand, 0, size),
	}
}

func (s *operandStack) push(v operand) {
	s.data = append(s.data, v)
	s.top++
}

func (s *operandStack) pop() operand {
	if s.top == 0 {
		log.Panic("stack underflow: no values to pop")
	}

	s.top--
	return s.data[s.top]
}
