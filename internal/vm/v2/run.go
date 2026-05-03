package vm

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) run(entryPoint uint32) error {
	m.frames.push(entryPoint, 0, m.src)

	for {
		opcode, _ := m.frames.current().readUint8(m.src)

		switch opcode {
		case emitter.StoreLocal:
			if err := m.storeLocal(); err != nil {
				return fmt.Errorf("error in StoreLocal: %w", err)
			}

		case emitter.LoadLocal:
			if err := m.loadLocal(); err != nil {
				return fmt.Errorf("error in LoadLocal: %w", err)
			}

		case emitter.PushString:
			if err := m.pushString(); err != nil {
				return fmt.Errorf("error in PushString: %w", err)
			}

		case emitter.PushInt:
			if err := m.pushInt(); err != nil {
				return fmt.Errorf("error in PushInt: %w", err)
			}

		case emitter.PushTrue:
			if err := m.pushTrue(); err != nil {
				return fmt.Errorf("error in PushTrue: %w", err)
			}

		case emitter.PushFalse:
			if err := m.pushFalse(); err != nil {
				return fmt.Errorf("error in PushFalse: %w", err)
			}

		case emitter.Call:
			if err := m.call(); err != nil {
				return fmt.Errorf("error in Call: %w", err)
			}

		case emitter.PushFn:
			if err := m.pushFn(); err != nil {
				return fmt.Errorf("error in PushFn: %w", err)
			}

		case emitter.PushNativeFn:
			if err := m.pushNativeFn(); err != nil {
				return fmt.Errorf("error in PushNativeFn: %w", err)
			}

		case emitter.JumpTo:
			offset, err := m.frames.current().readUint32(m.src)
			if err != nil {
				return fmt.Errorf("invalid jump offset operand: %w", err)
			}

			m.frames.current().moveTo(offset)

		case emitter.JumpIfFalse:
			offset, err := m.frames.current().readUint32(m.src)
			if err != nil {
				return fmt.Errorf("invalid jump offset operand: %w", err)
			}

			condition := m.operandStack.pop()

			if condition.ty != operandBool {
				return fmt.Errorf("expected boolean condition for JumpIfFalse, got %v", condition.ty)
			}

			isFalse := condition.boolValue == false
			if isFalse {
				m.frames.current().moveTo(offset)
			}

		case emitter.Pop:
			_ = m.operandStack.pop()

		case emitter.Add:
			if err := m.add(); err != nil {
				return fmt.Errorf("error in Add: %w", err)
			}

		case emitter.Sub:
			if err := m.sub(); err != nil {
				return fmt.Errorf("error in Sub: %w", err)
			}

		case emitter.Mul:
			if err := m.mul(); err != nil {
				return fmt.Errorf("error in Mul: %w", err)
			}

		case emitter.Div:
			if err := m.div(); err != nil {
				return fmt.Errorf("error in Div: %w", err)
			}

		case emitter.Eq:
			if err := m.eq(); err != nil {
				return fmt.Errorf("error in Eq: %w", err)
			}

		case emitter.Not:
			if err := m.not(); err != nil {
				return fmt.Errorf("error in Not: %w", err)
			}

		case emitter.Less:
			if err := m.less(); err != nil {
				return fmt.Errorf("error in Less: %w", err)
			}

		case emitter.LessEq:
			if err := m.lessEq(); err != nil {
				return fmt.Errorf("error in LessEq: %w", err)
			}

		case emitter.Return:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}
		}
	}
}
