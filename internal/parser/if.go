package parser

import "github.com/caiquetorres/lumi/internal/token"

type If struct {
	Condition Expr
	Then      *BlockExpr
}

func (p *Parser) parseIf() (*If, error) {
	_, err := p.lookahead().next().expect(token.If)
	if err != nil {
		return nil, err
	}

	condition, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	thenBlock, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &If{
		Condition: condition,
		Then:      thenBlock,
	}, nil
}
