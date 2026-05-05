package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type Loop struct {
	Body *Block

	span span.Span
}

func loop(body *Block, span span.Spanner) *Loop {
	return &Loop{
		Body: body,
		span: span.Span(),
	}
}
func (s *Loop) Span() span.Span {
	return s.span
}

var _ span.Spanner = (*Loop)(nil)

func (p *Parser) parseLoop() (*Loop, error) {
	loopTok, err := p.lookahead().next().expect(token.Loop)
	if err != nil {
		return nil, err
	}

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return loop(body, span.Merge(loopTok, body)), nil
}
