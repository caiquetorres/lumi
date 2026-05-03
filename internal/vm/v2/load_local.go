package vm

import (
	"encoding/binary"
	"fmt"
)

func (m *vm) loadLocal() error {
	offset, err := m.frames.current().readUint32(m.src)
	if err != nil {
		return err
	}

	data := m.readOffsetAt(offset)

	switch getTag(data) {
	case tagInt:
		val := decodeInt(data)
		op := intOperand(val)
		m.operandStack.push(op)

	case tagBool:
		val := decodeBool(data)
		op := boolOperand(val)
		m.operandStack.push(op)

	case tagString:
		strAddr := uint64(decodeString(data))
		op := stringOperand(strAddr)
		m.operandStack.push(op)

	case tagFn:
		fnIndex := uint32(decodeString(data))
		op := fnOperand(fnIndex)
		m.operandStack.push(op)

	default:
		return fmt.Errorf("unsupported tag for LoadLocal: %v", getTag(data))
	}

	return nil
}

func (m *vm) readOffsetAt(offset uint32) uint64 {
	pos := m.frames.current().sp + offset
	data := binary.LittleEndian.Uint64(m.locals[pos : pos+8])

	return data
}
