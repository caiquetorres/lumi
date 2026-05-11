package token

import "github.com/caiquetorres/lumi/internal/span"

type SymbolID int

type Token struct {
	s    span.Span
	kind Kind

	symbolID SymbolID
}

func New(id SymbolID, kind Kind, s span.Span) Token {
	return Token{
		kind:     kind,
		s:        s,
		symbolID: id,
	}
}

func (t Token) SymbolID() SymbolID {
	return t.symbolID
}

func (t Token) Kind() Kind {
	return t.kind
}

func (t Token) Span() span.Span {
	return t.s
}

var _ span.Spanner = Token{}
