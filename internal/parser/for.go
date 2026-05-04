package parser

import "github.com/caiquetorres/lumi/internal/token"

type ForStmt struct {
	Init  Stmt
	Cond  Expr
	Inc   Stmt
	Block *Block
}

func (p *Parser) parseFor() (*ForStmt, error) {
	_, err := p.lookahead().next().expect(token.For)
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

	if _, err := p.lookahead().next().expect(token.Semicolon); err != nil {
		return nil, err
	}

	var condExpr Expr
	if !p.lookahead().peek().is(token.Semicolon) {
		condExpr, err = p.parseExpr()
		if err != nil {
			return nil, err
		}
	}

	if _, err := p.lookahead().next().expect(token.Semicolon); err != nil {
		return nil, err
	}

	incStmt, err := p.parseStmt()
	if err != nil {
		return nil, err
	}

	block, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &ForStmt{
		Init:  initStmt,
		Cond:  condExpr,
		Inc:   incStmt,
		Block: block,
	}, nil
}
