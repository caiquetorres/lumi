package parser

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type CallExpr struct {
	Callee Expr
	Args   []Expr

	span span.Spanner
}

func callExpr(callee Expr, args []Expr, span span.Spanner) *CallExpr {
	return &CallExpr{
		Callee: callee,
		Args:   args,
		span:   span,
	}
}

func (c *CallExpr) Span() span.Span {
	return c.span.Span()
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
		tok, _ := p.lookahead().peek().get()
		return nil, &ParseError{
			Err:  fmt.Errorf("unexpected end of file: %w", ErrUnexpectedEOF),
			Span: tok.Span(),
		}
	}

	closeBraceTok, err := p.lookahead().next().expect(token.CloseParen)
	if err != nil {
		return nil, err
	}

	return callExpr(callee, args, span.Merge(callee, closeBraceTok)), nil
}
