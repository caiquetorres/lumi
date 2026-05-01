package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type LiteralKind int

const (
	LiteralString LiteralKind = iota + 1
	LiteralTrue
	LiteralFalse
	LiteralInt
)

type LiteralExpr struct {
	Kind  LiteralKind
	Value token.Token
}

func literalExpr(k LiteralKind, value token.Token) *LiteralExpr {
	return &LiteralExpr{
		Kind:  k,
		Value: value,
	}
}

func (l *LiteralExpr) expr() {}

var _ Expr = (*LiteralExpr)(nil)

func (p *Parser) isLiteral() bool {
	return p.lookahead().peek().isOneOf(
		token.String, token.True, token.False, token.Int,
	)
}

func (p *Parser) parseLiteral() (*LiteralExpr, error) {
	tok, err := p.lookahead().next().expectOneOf(
		token.String, token.True, token.False, token.Int,
	)

	if err != nil {
		return nil, err
	}

	switch tok.Kind() {
	case token.True:
		return literalExpr(LiteralTrue, tok), nil
	case token.False:
		return literalExpr(LiteralFalse, tok), nil
	case token.Int:
		return literalExpr(LiteralInt, tok), nil
	default:
		return literalExpr(LiteralString, tok), nil
	}
}
