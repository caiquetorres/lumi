package vm

import (
	"log"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) run(entryPoint uint32) error {
	m.frames.push(entryPoint, 0, m.src)

	for {
		opcode, _ := m.frames.current().readUint8(m.src)

		switch opcode {
		case emitter.PushInt:
			value, _ := m.frames.current().readUint32(m.src)

			m.operandStack.push(intOperand(int64(value)))

		case emitter.PushTrue:
			m.operandStack.push(boolOperand(true))

		case emitter.PushFalse:
			m.operandStack.push(boolOperand(false))

		case emitter.Add:
			right := m.operandStack.pop()
			left := m.operandStack.pop()

			if right.ty != operandInt || left.ty != operandInt {
				log.Panic("operands for Add must be integers")
			}

			result := left.intValue + right.intValue
			m.operandStack.push(intOperand(result))

		case emitter.Pop:
			_ = m.operandStack.pop()

		case emitter.Return:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}
		}
	}
}
