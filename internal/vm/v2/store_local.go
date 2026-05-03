package vm

import (
	"encoding/binary"
	"fmt"
)

func (m *vm) storeLocal() error {
	offset, err := m.frames.current().readUint32(m.src)
	if err != nil {
		return err
	}

	val := m.operandStack.pop()

	var data uint64
	switch val.ty {
	case operandInt:
		data = encodeInt(val.intValue)

	case operandBool:
		data = encodeBool(val.boolValue)

	case operandString:
		addr := int64(val.stringValue)
		data = encodeString(addr)

	case operandFn:
		fnIdx := int64(val.fnValue)
		data = encodeFn(fnIdx)

	default:
		return fmt.Errorf("unsupported operand type for StoreLocal: %v", val.ty)
	}

	m.writeOffsetAt(offset, data)

	return nil
}

func (m *vm) writeOffsetAt(offset uint32, value uint64) {
	pos := m.frames.current().sp + offset
	binary.LittleEndian.PutUint64(m.locals[pos:pos+8], value)
	m.frames.current().tmp = pos + 8
}
