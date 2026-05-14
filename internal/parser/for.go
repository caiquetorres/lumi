package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type For struct {
	Init Stmt
	Cond Expr
	Inc  Stmt
	Body *Block

	span span.Span
}

func forStmt(init Stmt, cond Expr, inc Stmt, body *Block, span span.Span) *For {
	return &For{
		Init: init,
		Cond: cond,
		Inc:  inc,
		Body: body,
		span: span,
	}
}

func (s *For) Span() span.Span {
	return s.span
}

func (p *Parser) parseFor() (*For, error) {
	forTok, err := p.lookahead().next().expect(token.For)
	if err != nil {
		return nil, err
	}

	var initStmt Stmt
	if !p.lookahead().peek().is(token.Semicolon) {
		initStmt, err = p.parseStmt()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.lookahead().next().expect(token.Semicolon)
	if err != nil {
		return nil, err
	}

	var condExpr Expr
	if !p.lookahead().peek().is(token.Semicolon) {
		condExpr, err = p.parseExpr()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.lookahead().next().expect(token.Semicolon)
	if err != nil {
		return nil, err
	}

	var incStmt Stmt
	if !p.lookahead().peek().is(token.OpenBrace) {
		incStmt, err = p.parseStmt()
		if err != nil {
			return nil, err
		}
	}

	block, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return forStmt(
		initStmt, condExpr, incStmt, block,
		span.Merge(forTok, block),
	), nil
}
