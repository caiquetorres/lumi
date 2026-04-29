package parser

import "github.com/caiquetorres/lumi/internal/token"

type If struct {
	Condition Expr
	Then      *Block
	Else      *Block
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

	var elseBlock *Block
	if p.lookahead().peek().is(token.Else) {
		p.lookahead().next() // consume 'else'

		if p.lookahead().peek().is(token.If) {
			elseIf, err := p.parseIf()
			if err != nil {
				return nil, err
			}

			elseBlock = &Block{
				Stmts: []Stmt{elseIf},
			}
		} else {
			elseBlock, err = p.parseBlock()
			if err != nil {
				return nil, err
			}
		}
	}

	return &If{
		Condition: condition,
		Then:      thenBlock,
		Else:      elseBlock,
	}, nil
}
