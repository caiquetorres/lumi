package emitter

type loop struct {
	start     uint32
	condStart uint32
	end       []uint32
}

type loopStack struct {
	data []loop
}

func newLoopStack() *loopStack {
	return &loopStack{
		data: make([]loop, 0),
	}
}

func (s *loopStack) push(loop loop) {
	s.data = append(s.data, loop)
}

func (s *loopStack) pop() (loop, bool) {
	if len(s.data) == 0 {
		return loop{}, false
	}

	n := len(s.data) - 1
	loop := s.data[n]
	s.data = s.data[:n]
	return loop, true
}

func (s *loopStack) top() (*loop, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	n := len(s.data) - 1
	return &s.data[n], true
}
