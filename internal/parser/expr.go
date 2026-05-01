package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Expr interface {
	expr() // marker method
}

func (p *Parser) parseExpr() (Expr, error) {
	unit, err := p.parseUnit()
	if err != nil {
		return nil, err
	}

	if p.lookahead().peek().is(token.OpenParen) {
		return p.parseCallExpr(unit)
	}

	return unit, nil
}

func (p *Parser) parseUnit() (Expr, error) {
	switch {
	case p.isLiteral():
		return p.parseLiteral()
	case p.lookahead().peek().is(token.Identifier):
		return p.parseIdentifier()
	default:
		_, err := p.lookahead().peek().expectOneOf(token.String)
		return nil, err
	}
}
