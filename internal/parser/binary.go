package parser

import (
	"slices"

	"github.com/caiquetorres/lumi/internal/token"
)

type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (b *BinaryExpr) expr() {}

var _ Expr = (*BinaryExpr)(nil)

func binaryExpr(left Expr, operator token.Token, right Expr) *BinaryExpr {
	return &BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (p *Parser) parseBinaryExpr(parentPrec int) (Expr, error) {
	left, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	tok, err := p.lookahead().peek().get()
	if err != nil {
		return nil, err
	}

	if !isBinaryOperator(tok) {
		return left, nil
	}

	prec := precedence(tok.Kind())
	for prec > parentPrec {
		opTok, err := p.lookahead().next().get()
		if err != nil {
			return nil, err
		}

		right, err := p.parseBinaryExpr(prec)
		if err != nil {
			return nil, err
		}

		left = binaryExpr(left, opTok, right)

		tok, err := p.lookahead().peek().get()
		if err != nil {
			return nil, err
		}

		if !isBinaryOperator(tok) {
			break
		}

		prec = precedence(tok.Kind())
	}

	return left, nil
}

func (p *Parser) parseTerm() (Expr, error) {
	expr, err := p.parseUnit()
	if err != nil {
		return nil, err
	}

	if p.lookahead().peek().is(token.OpenParen) {
		return p.parseCallExpr(expr)
	}

	return expr, nil
}

// func isUnaryOperator(tok token.Token) bool {
// 	switch tok.Kind() {
// 	case token.Bang, token.Plus, token.Minus:
// 		return true
// 	default:
// 		return false
// 	}
// }

func isBinaryOperator(tok token.Token) bool {
	switch tok.Kind() {
	case token.Plus, token.Minus, token.Star, token.Slash,
		token.Equal, token.EqualEqual, token.BangEqual,
		token.Less, token.LessEqual, token.Greater,
		token.GreaterEqual, token.PlusEqual:
		return true
	default:
		return false
	}
}

func precedence(kind token.Kind) int {
	prec := [][]token.Kind{
		// higher precedence
		{token.Star, token.Slash},
		{token.Plus, token.Minus},
		{token.Less, token.LessEqual, token.Greater, token.GreaterEqual},
		{token.EqualEqual, token.BangEqual},
		{token.Equal, token.PlusEqual},
		// lower precedence
	}

	for i, group := range prec {
		if slices.Contains(group, kind) {
			return len(prec) - 1 - i
		}
	}

	return -1
}
