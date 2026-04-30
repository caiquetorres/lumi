package parser

import "github.com/caiquetorres/lumi/internal/token"

type While struct {
	Condition Expr
	Body      *Block
}

func (p *Parser) parseWhile() (*While, error) {
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

	return &While{
		Condition: condition,
		Body:      body,
	}, nil
}
