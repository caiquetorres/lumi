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

	if p.lookahead().peek().is(token.OpenParen) {
		return p.parseCallExpr(unit)
	}

	return unit, nil
}

type LiteralExpr struct {
	Kind  LiteralKind
	Value token.Token
}

func (p *Parser) parseLiteral() (*LiteralExpr, error) {
	tok, err := p.lookahead().next().expect(token.String)
	if err != nil {
		return nil, err
	}

	return &LiteralExpr{
		Kind:  LiteralString,
		Value: tok,
	}, nil
}

type IdentifierExpr struct {
	Name token.Token
}

func (p *Parser) parseIdentifier() (*IdentifierExpr, error) {
	tok, err := p.lookahead().next().expect(token.Identifier)
	if err != nil {
		return nil, err
	}

	return &IdentifierExpr{
		Name: tok,
	}, nil
}

func (p *Parser) parseUnit() (Expr, error) {
	switch {
	case p.lookahead().peek().is(token.String):
		return p.parseLiteral()
	case p.lookahead().peek().is(token.Identifier):
		return p.parseIdentifier()
	case p.lookahead().peek().is(token.OpenBrace):
		return p.parseBlock()
	default:
		return p.lookahead().peek().expectOneOf(token.String)
	}
}

type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func (p *Parser) parseCallExpr(callee Expr) (Expr, error) {
	_, err := p.lookahead().next().expect(token.OpenParen)
	if err != nil {
		return nil, err
	}

	var args []Expr
	for !p.lookahead().peek().isOneOf(token.CloseParen, token.EOF) {
		arg, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		_, err = p.lookahead().peek().expectOneOf(token.Comma, token.CloseParen)
		if err != nil {
			return nil, err
		}

		if p.lookahead().peek().is(token.Comma) {
			p.bump() // close paren
		}
	}

	if p.lookahead().peek().is(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	p.bump() // close brace

	return &CallExpr{
		Callee: callee,
		Args:   args,
	}, nil
}
