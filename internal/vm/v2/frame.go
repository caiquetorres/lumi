package vm

import (
	"encoding/binary"
	"io"
	"log"
)

const (
	uint8Size  = 1
	uint32Size = 4
)

type frame struct {
	sp uint32 // stack pointer
	ip uint32 // instruction pointer and bytecode
}

func (f *frame) moveTo(offset uint32) {
	f.ip = offset
}

func (f *frame) hasMore(data []byte) bool {
	return f.ip < uint32(len(data))
}

func (f *frame) readUint8(data []byte) (uint8, error) {
	if f.ip+uint8Size > uint32(len(data)) {
		return 0, io.EOF
	}

	value := data[f.ip]
	f.ip += uint8Size

	return value, nil
}

func (f *frame) readUint32(data []byte) (uint32, error) {
	if f.ip+uint32Size > uint32(len(data)) {
		return 0, io.EOF
	}

	value := binary.BigEndian.Uint32(data[f.ip : f.ip+uint32Size])

	f.ip += uint32Size

	return value, nil
}

const MAX_STACK_SIZE = 1024

type frames struct {
	data []frame
}

func newFrames() *frames {
	return &frames{
		data: make([]frame, 0, MAX_STACK_SIZE),
	}
}

func (f *frames) reset() {
	f.data = f.data[:0]
}

func (f *frames) current() *frame {
	if len(f.data) == 0 {
		log.Panic("no frames available: cannot get current frame")
	}

	top := len(f.data) - 1
	return &f.data[top]
}

func (f *frames) isEmpty() bool {
	return len(f.data) == 0
}

func (f *frames) pop() {
	if len(f.data) == 0 {
		log.Panic("stack underflow: no frames to pop")
	}

	f.data = f.data[:len(f.data)-1]
}

func (f *frames) push(ip, sp uint32, data []byte) {
	if len(f.data) >= MAX_STACK_SIZE {
		log.Panic("stack overflow: too many frames")
	}

	f.data = append(f.data, frame{
		ip: ip,
		sp: sp,
	})
}
