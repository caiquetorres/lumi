package emitter

type jumpStack struct {
	data []uint32
}

func newJumpStack() *jumpStack {
	return &jumpStack{
		data: make([]uint32, 0),
	}
}

func (s *jumpStack) push(offset uint32) {
	s.data = append(s.data, offset)
}

func (s *jumpStack) pop() (offset uint32, ok bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	n := len(s.data) - 1
	offset = s.data[n]
	s.data = s.data[:n]
	return offset, true
}

func (s *jumpStack) top() (offset uint32, ok bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	n := len(s.data) - 1
	return s.data[n], true
}
