package parser

import "github.com/caiquetorres/lumi/internal/token"

type ReturnStmt struct {
	Expr Expr
}

func returnStmt(expr Expr) *ReturnStmt {
	return &ReturnStmt{
		Expr: expr,
	}
}

func (p *Parser) parseReturn() (*ReturnStmt, error) {
	_, err := p.lookahead().next().expect(token.Return)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return returnStmt(expr), nil
}
