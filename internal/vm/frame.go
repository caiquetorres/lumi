package vm

import "log"

const MAX_STACK_SIZE = 1024

type frame struct {
	cursor
	prevSymbolTable *symbolTable
}

type frames struct {
	data []frame
}

func (f *frames) current() *frame {
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

func (f *frames) push(ptr uint32, data []byte, savedTable *symbolTable) {
	if len(f.data) >= MAX_STACK_SIZE {
		log.Panic("stack overflow: too many frames")
	}

	f.data = append(f.data, frame{
		cursor: cursor{
			pc:   ptr,
			data: data,
		},
		prevSymbolTable: savedTable,
	})
}
