package token

import "github.com/caiquetorres/lumi/internal/span"

type SymbolID int

type Token struct {
	kind Kind
	s    span.Span

	symbolID  SymbolID
	hasSymbol bool
}

func New(id int, kind Kind, s span.Span) Token {
	return Token{
		kind: kind,
		s:    s,
	}
}

func NewWithSymbol(id int, kind Kind, s span.Span) Token {
	return Token{
		kind:      kind,
		s:         s,
		symbolID:  SymbolID(id),
		hasSymbol: true,
	}
}

func (t Token) SymbolID() SymbolID {
	if !t.hasSymbol {
		return -1
	}

	return t.symbolID
}

func (t Token) Kind() Kind {
	return t.kind
}

func (t Token) Span() span.Span {
	return t.s
}

var _ span.Spanner = Token{}
