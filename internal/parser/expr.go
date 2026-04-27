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
	LiteralTrue
	LiteralFalse
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

func (p *Parser) isLiteral() bool {
	return p.
		lookahead().
		peek().
		isOneOf(token.String, token.True, token.False)
}

func (p *Parser) parseLiteral() (*LiteralExpr, error) {
	tok, err := p.
		lookahead().
		next().
		expectOneOf(token.String, token.True, token.False)

	if err != nil {
		return nil, err
	}

	switch tok.Kind() {
	case token.True:
		return &LiteralExpr{Kind: LiteralTrue, Value: tok}, nil
	case token.False:
		return &LiteralExpr{Kind: LiteralFalse, Value: tok}, nil
	default:
		return &LiteralExpr{Kind: LiteralString, Value: tok}, nil
	}
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
	case p.isLiteral():
		return p.parseLiteral()
	case p.lookahead().peek().is(token.Identifier):
		return p.parseIdentifier()
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
