package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func callExpr(callee Expr, args []Expr) *CallExpr {
	return &CallExpr{
		Callee: callee,
		Args:   args,
	}
}

var _ Expr = (*CallExpr)(nil)

func (c *CallExpr) expr() {}

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

	return callExpr(callee, args), nil
}
