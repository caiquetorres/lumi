package vm

import (
	"encoding/binary"
	"errors"
)

type heap struct {
	memory []byte
	offset int64
}

var (
	ErrOutOfMemory    = errors.New("out of memory")
	ErrInvalidAddress = errors.New("invalid memory address")
)

func newHeap(size int) *heap {
	return &heap{
		memory: make([]byte, size),
		offset: 0,
	}
}

func (h *heap) alloc(size int) (int64, error) {
	if int(h.offset)+size > len(h.memory) {
		return 0, ErrOutOfMemory
	}

	addr := h.offset
	h.offset += int64(size)
	return addr, nil
}

func (h *heap) writeInt32(addr int64, value int32) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(value))

	return h.write(addr, buf[:])
}

func (h *heap) read(addr int, size int) ([]byte, error) {
	if addr < 0 || addr+size > len(h.memory) {
		return nil, ErrInvalidAddress
	}

	return h.memory[addr : addr+size], nil
}

func (h *heap) write(addr int64, data []byte) error {
	if addr < 0 || addr+int64(len(data)) > int64(len(h.memory)) {
		return ErrInvalidAddress
	}

	copy(h.memory[addr:], data)
	return nil
}
