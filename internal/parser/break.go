package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type BreakStmt struct {
	span span.Span
}

func breakStmt(span span.Spanner) *BreakStmt {
	return &BreakStmt{
		span: span.Span(),
	}
}

func (s *BreakStmt) Span() span.Span {
	return s.span
}

func (p *Parser) parseBreak() (*BreakStmt, error) {
	breakTok, err := p.lookahead().next().expect(token.Break)
	if err != nil {
		return nil, err
	}

	return breakStmt(breakTok), nil
}
