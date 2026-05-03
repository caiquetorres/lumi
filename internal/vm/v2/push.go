package vm

import "fmt"

func (m *vm) pushTrue() error {
	op := boolOperand(true)
	m.operandStack.push(op)

	return nil
}

func (m *vm) pushFalse() error {
	op := boolOperand(false)
	m.operandStack.push(op)

	return nil
}

func (m *vm) pushInt() error {
	val, err := m.frames.current().readUint32(m.src)
	if err != nil {
		return err
	}

	op := intOperand(int64(val))
	m.operandStack.push(op)

	return nil
}

func (m *vm) pushFn() error {
	fnIndex, err := m.frames.current().readUint32(m.src)
	if err != nil {
		return err
	}

	fnOffset := m.fnTable[fnIndex]

	op := fnOperand(fnOffset)
	m.operandStack.push(op)

	return nil
}

func (m *vm) pushNativeFn() error {
	fnNameIndex, err := m.frames.current().readUint32(m.src)
	if err != nil {
		return err
	}

	op := nativeFnOperand(fnNameIndex)
	m.operandStack.push(op)

	return nil
}

func (m *vm) pushString() error {
	constIdx, err := m.frames.current().readUint32(m.src)
	if err != nil {
		return err
	}

	constStr, exists := m.pool.GetConstant(constIdx)
	if !exists {
		return fmt.Errorf("constant not found")
	}

	str := []byte(constStr.(string))
	strObj := heapObject{
		tag:  tagString,
		size: len(str),
		data: str,
	}

	addr, err := m.heap.allocAndWriteObject(strObj)
	if err != nil {
		return err
	}

	srtAddr := encodeString(addr)
	op := stringOperand(srtAddr)
	m.operandStack.push(op)

	return nil
}
