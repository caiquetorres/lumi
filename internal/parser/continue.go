package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type Continue struct {
	span span.Span
}

func continueStmt(span span.Spanner) *Continue {
	return &Continue{
		span: span.Span(),
	}
}

func (s *Continue) Span() span.Span {
	return s.span
}

func (p *Parser) parseContinue() (*Continue, error) {
	continueTok, err := p.lookahead().next().expect(token.Continue)
	if err != nil {
		return nil, err
	}

	return continueStmt(continueTok), nil
}
