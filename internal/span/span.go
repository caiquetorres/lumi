package span

type Spanner interface {
	Span() Span
}

type Span struct {
	start, end int
}

func (s Span) Span() Span {
	return s
}

func New(start, end int) Span {
	return Span{
		start: start,
		end:   end,
	}
}

func (s Span) Start() int {
	return s.start
}

func (s Span) End() int {
	return s.end
}

func (s Span) Len() int {
	return s.end - s.start
}

var _ Spanner = Span{}
