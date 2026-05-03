package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) run(entryPoint uint32) error {
	m.frames.push(entryPoint, 0, m.src)

	for {
		opcode, _ := m.frames.current().readUint8(m.src)

		switch opcode {
		case emitter.StoreLocal:
			offset, _ := m.frames.current().readUint32(m.src)
			val := m.operandStack.pop()

			var data uint64
			switch val.ty {
			case operandInt:
				data = encodeInt(val.intValue)
			case operandBool:
				data = encodeBool(val.boolValue)
			case operandString:
				data = encodeString(int64(val.strValue))
			default:
				return fmt.Errorf("unsupported operand type for LoadLocal: %v", val.ty)
			}

			pos := m.frames.current().sp + offset
			binary.LittleEndian.PutUint64(m.locals[pos:pos+8], data)
			m.frames.current().tmp = pos

		case emitter.LoadLocal:
			offset, _ := m.frames.current().readUint32(m.src)

			pos := m.frames.current().sp + offset
			data := binary.LittleEndian.Uint64(m.locals[pos : pos+8])

			switch getTag(data) {
			case tagInt:
				m.operandStack.push(intOperand(decodeInt(data)))
			case tagBool:
				m.operandStack.push(boolOperand(decodeBool(data)))
			case tagString:
				strAddr := decodeString(data)
				m.operandStack.push(stringOperand(uint64(strAddr)))
			default:
				return fmt.Errorf("unsupported tag for LoadLocal: %v", getTag(data))
			}

		case emitter.PushString:
			constIdx, _ := m.frames.current().readUint32(m.src)
			constStr, _ := m.pool.GetConstant(constIdx)

			str := []byte(constStr.(string))
			strObj := heapObject{
				tag:  tagString,
				size: len(str),
				data: str,
			}

			addr, _ := m.heap.allocAndWriteObject(strObj)
			m.operandStack.push(stringOperand(encodeString(addr)))

		case emitter.PushInt:
			value, _ := m.frames.current().readUint32(m.src)

			m.operandStack.push(intOperand(int64(value)))

		case emitter.PushTrue:
			m.operandStack.push(boolOperand(true))

		case emitter.PushFalse:
			m.operandStack.push(boolOperand(false))

		case emitter.Call:
			val := m.operandStack.pop()

			if val.ty != operandFn && val.ty != operandNativeFn {
				return fmt.Errorf("expected function operand for Call, got %v", val.ty)
			}

			arity, _ := m.frames.current().readUint8(m.src)

			if val.ty == operandFn {
				tmp := m.frames.current().tmp
				m.frames.push(val.fnValue, tmp, m.src)
			} else {
				fnNameIndex := val.fnValue
				fnNameConst, _ := m.pool.GetConstant(fnNameIndex)

				operands := make([]operand, arity)
				for i := int(arity) - 1; i >= 0; i-- {
					operands[i] = m.operandStack.pop()
				}

				res, err := m.nativeFnTable[fnNameConst.(string)](m, operands...)
				if err != nil {
					return fmt.Errorf("error calling native function: %w", err)
				}

				m.operandStack.push(res)
			}

		case emitter.PushFn:
			fnIndex, _ := m.frames.current().readUint32(m.src)
			fnOffset := m.fnTable[fnIndex]
			m.operandStack.push(fnOperand(fnOffset))

		case emitter.PushNativeFn:
			fnNameIndex, _ := m.frames.current().readUint32(m.src)
			m.operandStack.push(nativeFnOperand(fnNameIndex))

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
			right := m.operandStack.pop()
			left := m.operandStack.pop()

			if right.ty != operandInt || left.ty != operandInt {
				return fmt.Errorf("operands for Add must be integers")
			}

			result := left.intValue + right.intValue
			m.operandStack.push(intOperand(result))

		case emitter.Sub:
			right := m.operandStack.pop()
			left := m.operandStack.pop()

			if right.ty != operandInt || left.ty != operandInt {
				return fmt.Errorf("operands for Sub must be integers")
			}

			result := left.intValue - right.intValue
			m.operandStack.push(intOperand(result))

		case emitter.Mul:
			right := m.operandStack.pop()
			left := m.operandStack.pop()

			if right.ty != operandInt || left.ty != operandInt {
				return fmt.Errorf("operands for Mul must be integers")
			}

			result := left.intValue * right.intValue
			m.operandStack.push(intOperand(result))

		case emitter.Div:
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

		case emitter.Eq:
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
			default:
				return fmt.Errorf("unsupported operand type for Eq: %v", right.ty)
			}

			m.operandStack.push(boolOperand(result))

		case emitter.Not:
			value := m.operandStack.pop()

			if value.ty != operandBool {
				return fmt.Errorf("operand for Not must be a boolean")
			}

			m.operandStack.push(boolOperand(!value.boolValue))

		case emitter.Less:
			right := m.operandStack.pop()
			left := m.operandStack.pop()

			if right.ty != operandInt || left.ty != operandInt {
				return fmt.Errorf("operands for Less must be integers")
			}

			result := left.intValue < right.intValue
			m.operandStack.push(boolOperand(result))

		case emitter.LessEq:
			right := m.operandStack.pop()
			left := m.operandStack.pop()

			if right.ty != operandInt || left.ty != operandInt {
				return fmt.Errorf("operands for LessEq must be integers")
			}

			result := left.intValue <= right.intValue
			m.operandStack.push(boolOperand(result))

		case emitter.Return:
			m.frames.pop()

			if m.frames.isEmpty() {
				return nil
			}
		}
	}
}
