package vm

const MAX_STACK_SIZE = 1024

type frames struct {
	data []cursor
}

func (f *frames) current() *cursor {
	if len(f.data) == 0 {
		panic("no frames available: cannot get current frame")
	}

	return &f.data[len(f.data)-1]
}

func (f *frames) incCurrentPtr(offset uint32) {
	if len(f.data) == 0 {
		return
	}

	f.data[len(f.data)-1].move(offset)
}

func (f *frames) isEmpty() bool {
	return len(f.data) == 0
}

func (f *frames) pop() {
	if len(f.data) == 0 {
		panic("stack underflow: no frames to pop")
	}

	f.data = f.data[:len(f.data)-1]
}

func (f *frames) push(ptr uint32, data []byte) {
	if len(f.data) >= MAX_STACK_SIZE {
		panic("stack overflow: too many frames")
	}

	f.data = append(f.data, cursor{
		pc:   ptr,
		data: data,
	})
}
