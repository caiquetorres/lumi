package token

import "github.com/caiquetorres/lumi/internal/span"

type Token struct {
	kind Kind
	s    span.Span
}

func New(kind Kind, s span.Span) Token {
	return Token{
		kind: kind,
		s:    s,
	}
}

func (t Token) Kind() Kind {
	return t.kind
}

func (t Token) Span() span.Span {
	return t.s
}

var _ span.Spanner = Token{}
