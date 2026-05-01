package parser

import "github.com/caiquetorres/lumi/internal/token"

type WhileStmt struct {
	Condition Expr
	Body      *Block
}

func whileStmt(condition Expr, body *Block) *WhileStmt {
	return &WhileStmt{
		Condition: condition,
		Body:      body,
	}
}

func (p *Parser) parseWhile() (*WhileStmt, error) {
	_, err := p.lookahead().next().expect(token.While)
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

	return whileStmt(condition, body), nil
}
