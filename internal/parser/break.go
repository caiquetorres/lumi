package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type Break struct {
	span span.Span
}

func breakStmt(span span.Spanner) *Break {
	return &Break{
		span: span.Span(),
	}
}

func (s *Break) Span() span.Span {
	return s.span
}

func (p *Parser) parseBreak() (*Break, error) {
	breakTok, err := p.lookahead().next().expect(token.Break)
	if err != nil {
		return nil, err
	}

	return breakStmt(breakTok), nil
}
