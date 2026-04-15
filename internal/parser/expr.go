package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type (
	Expr any
	Stmt Expr
)

type LiteralKind int

const (
	LiteralString LiteralKind = iota + 1
)

func (p *Parser) parseExpr() (Expr, error) {
	unit, err := p.parseUnit()
	if err != nil {
		return nil, err
	}

	if p.peekIs(token.OpenParen) {
		return p.parseCallExpr(unit)
	}

	return unit, nil
}

type LiteralExpr struct {
	Kind  LiteralKind
	Value token.Token
}

type IdentifierExpr struct {
	Name token.Token
}

func (p *Parser) parseUnit() (Expr, error) {
	switch {
	case p.peekIs(token.String):
		tok, _ := p.expect(token.String)

		return &LiteralExpr{
			Kind:  LiteralString,
			Value: tok,
		}, nil
	case p.peekIs(token.Identifier):
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

type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func (p *Parser) parseCallExpr(callee Expr) (Expr, error) {
	if _, err := p.expect(token.OpenParen); err != nil {
		return nil, err
	}

	var args []Expr
	for !p.peekIs(token.CloseParen) {
		arg, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		if !p.peekIsOneOf(token.Comma, token.CloseParen) {
			_, err := p.expectOneOf(token.Comma, token.CloseParen)
			return nil, err
		}

		if p.peekIs(token.Comma) {
			_, _ = p.next()
		}
	}

	if p.peekIs(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	_, _ = p.next() // close brace

	return &CallExpr{
		Callee: callee,
		Args:   args,
	}, nil
}
