package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type ContinueStmt struct {
	span span.Span
}

func continueStmt(span span.Spanner) *ContinueStmt {
	return &ContinueStmt{
		span: span.Span(),
	}
}

func (s *ContinueStmt) Span() span.Span {
	return s.span
}

func (p *Parser) parseContinue() (*ContinueStmt, error) {
	continueTok, err := p.lookahead().next().expect(token.Continue)
	if err != nil {
		return nil, err
	}

	return continueStmt(continueTok), nil
}
