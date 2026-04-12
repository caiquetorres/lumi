package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Expr any

type LiteralKind int

const (
	LiteralString LiteralKind = iota + 1
)

type LiteralExpr struct {
	Kind  LiteralKind
	Value token.Token
}

type CallExpr struct {
	Callee Expr
}

type IdentifierExpr struct {
	Name token.Token
}

func (p *Parser) parseExpression() (Expr, error) {
	unit, err := p.parseUnit()
	if err != nil {
		return nil, err
	}

	if p.is(token.OpenParen) {
		return p.parseCallExpr(unit)
	}

	return unit, nil
}

func (p *Parser) parseUnit() (Expr, error) {
	switch {
	case p.is(token.String):
		tok, _ := p.expect(token.String)

		return &LiteralExpr{
			Kind:  LiteralString,
			Value: tok,
		}, nil
	case p.is(token.Identifier):
		tok, _ := p.expect(token.Identifier)

		return &IdentifierExpr{
			Name: tok,
		}, nil
	default:
		if err := p.err(); err != nil {
			return nil, err
		}

		return p.expectOneOf(token.String)
	}
}

func (p *Parser) parseCallExpr(callee Expr) (Expr, error) {
	if _, err := p.expect(token.OpenParen); err != nil {
		return nil, err
	}

	if _, err := p.expect(token.CloseParen); err != nil {
		return nil, err
	}

	return &CallExpr{
		Callee: callee,
	}, nil
}
