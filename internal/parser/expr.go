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

	if p.peek().is(token.OpenParen) {
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
	case p.peek().is(token.String):
		tok, _ := p.next().get()

		return &LiteralExpr{
			Kind:  LiteralString,
			Value: tok,
		}, nil
	case p.peek().is(token.Identifier):
		tok, _ := p.next().get()

		return &IdentifierExpr{
			Name: tok,
		}, nil
	default:
		return p.peek().expectOneOf(token.String)
	}
}

type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func (p *Parser) parseCallExpr(callee Expr) (Expr, error) {
	_, err := p.next().expect(token.OpenParen)
	if err != nil {
		return nil, err
	}

	var args []Expr
	for !p.peek().isOneOf(token.CloseParen, token.EOF) {
		arg, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		if !p.peek().isOneOf(token.Comma, token.CloseParen) {
			_, err := p.peek().expectOneOf(token.Comma, token.CloseParen)
			return nil, err
		}

		p.maybeNext(token.Comma)
	}

	if p.peek().is(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	p.bump() // close brace

	return &CallExpr{Callee: callee, Args: args}, nil
}
