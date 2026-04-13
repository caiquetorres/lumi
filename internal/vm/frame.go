package vm

const MAX_STACK_SIZE = 1024

type frame struct {
	ptr uint32
}

type frames struct {
	data []frame
}

func newFrames(entry uint32) frames {
	return frames{
		data: []frame{{ptr: entry}},
	}
}

func (f *frames) current() *frame {
	if len(f.data) == 0 {
		return nil
	}

	return &f.data[len(f.data)-1]
}

func (f *frames) incrementCurrentPtr(offset uint32) {
	if len(f.data) == 0 {
		return
	}

	f.data[len(f.data)-1].ptr += offset
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

func (f *frames) push(ptr uint32) {
	if len(f.data) >= MAX_STACK_SIZE {
		panic("stack overflow: too many frames")
	}

	f.data = append(f.data, frame{ptr: ptr})
}
