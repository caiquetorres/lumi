package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type Return struct {
	Expr Expr

	span span.Span
}

func returnStmt(expr Expr, span span.Spanner) *Return {
	return &Return{
		Expr: expr,
		span: span.Span(),
	}
}

func (s *Return) Span() span.Span {
	return s.span
}

func (p *Parser) parseReturn() (*Return, error) {
	returnTok, err := p.lookahead().next().expect(token.Return)
	if err != nil {
		return nil, err
	}

	if p.lookahead().peek().is(token.NewLine) {
		return returnStmt(nil, returnTok.Span()), nil
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return returnStmt(expr, span.Merge(returnTok, expr)), nil
}
