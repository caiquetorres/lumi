package vm

import (
	"encoding/binary"
	"errors"
)

type heapObject struct {
	tag  tag
	size int
	data []byte
}

type heap struct {
	memory []byte
	offset int
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

func (h *heap) alloc(size int) (int, error) {
	if h.offset+size > len(h.memory) {
		return 0, ErrOutOfMemory
	}

	addr := h.offset
	h.offset += size
	return addr, nil
}

func (h *heap) allocAndWriteObject(obj heapObject) (int, error) {
	totalSize := 1 + 4 + obj.size // 1 byte for tag, 4 bytes for size, rest for data

	addr, err := h.alloc(totalSize)
	if err != nil {
		return 0, err
	}

	h.memory[addr] = byte(obj.tag)
	binary.BigEndian.PutUint32(h.memory[addr+1:addr+5], uint32(obj.size))
	copy(h.memory[addr+5:], obj.data)

	return addr, nil
}

func (h *heap) writeInt32(addr int, value int32) error {
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

func (h *heap) readObject(addr int) (heapObject, error) {
	if addr < 0 || addr+5 > len(h.memory) {
		return heapObject{}, ErrInvalidAddress
	}

	tag := tag(h.memory[addr])
	size := int(binary.BigEndian.Uint32(h.memory[addr+1 : addr+5]))

	if addr+5+size > len(h.memory) {
		return heapObject{}, ErrInvalidAddress
	}

	data := h.memory[addr+5 : addr+5+size]
	return heapObject{
		tag:  tag,
		size: size,
		data: data,
	}, nil
}

func (h *heap) write(addr int, data []byte) error {
	if addr < 0 || addr+len(data) > len(h.memory) {
		return ErrInvalidAddress
	}

	copy(h.memory[addr:], data)
	return nil
}
