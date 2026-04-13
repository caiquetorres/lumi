package span

type Spanner interface {
	Span() Span
}

type Span struct {
	start, end uint32
}

func New(start, end uint32) Span {
	return Span{
		start: start,
		end:   end,
	}
}

func (s Span) Span() Span {
	return s
}

func (s Span) Start() uint32 {
	return s.start
}

func (s Span) End() uint32 {
	return s.end
}

func (s Span) Len() uint32 {
	return s.end - s.start
}

var _ Spanner = Span{}
