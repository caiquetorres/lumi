package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type WhileStmt struct {
	Condition Expr
	Body      *Block

	span span.Span
}

func whileStmt(condition Expr, body *Block, span span.Spanner) *WhileStmt {
	return &WhileStmt{
		Condition: condition,
		Body:      body,
		span:      span.Span(),
	}
}

func (s *WhileStmt) Span() span.Span {
	return s.span
}

func (p *Parser) parseWhile() (*WhileStmt, error) {
	whileTok, err := p.lookahead().next().expect(token.While)
	if err != nil {
		return nil, err
	}

	condition, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return whileStmt(condition, body, span.Merge(whileTok, body)), nil
}
