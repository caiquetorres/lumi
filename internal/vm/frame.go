package vm

import "log"

const MAX_STACK_SIZE = 1024

type frames struct {
	data []cursor
}

func (f *frames) current() *cursor {
	if len(f.data) == 0 {
		log.Panic("no frames available: cannot get current frame")
	}

	return &f.data[len(f.data)-1]
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

func (f *frames) push(ptr uint32, data []byte) {
	if len(f.data) >= MAX_STACK_SIZE {
		log.Panic("stack overflow: too many frames")
	}

	f.data = append(f.data, cursor{
		pc:   ptr,
		data: data,
	})
}
