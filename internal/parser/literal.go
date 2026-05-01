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

func literal(k LiteralKind, value token.Token) *LiteralExpr {
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
		return literal(LiteralTrue, tok), nil
	case token.False:
		return literal(LiteralFalse, tok), nil
	case token.Int:
		return literal(LiteralInt, tok), nil
	default:
		return literal(LiteralString, tok), nil
	}
}
