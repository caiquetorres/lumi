package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Expression any

type LiteralKind int

const (
	LiteralString LiteralKind = iota + 1
)

type LiteralExpr struct {
	Kind  LiteralKind
	Value token.Token
}

func (p *Parser) parseExpression() (Expression, error) {
	switch {
	case p.is(token.String):
		tok, _ := p.expect(token.String)

		return &LiteralExpr{
			Kind:  LiteralString,
			Value: tok,
		}, nil
	default:
		if err := p.err(); err != nil {
			return nil, err
		}

		return p.expectOneOf(token.String)
	}
}
