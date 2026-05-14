package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type If struct {
	Condition Expr
	Then      *Block
	Else      *Block

	span span.Span
}

func ifStmt(
	condition Expr,
	thenBlock, elseBlock *Block,
	span span.Spanner,
) *If {
	return &If{
		Condition: condition,
		Then:      thenBlock,
		Else:      elseBlock,
		span:      span.Span(),
	}
}

func (s *If) Span() span.Span {
	return s.span
}

func (p *Parser) parseIf() (*If, error) {
	ifTok, err := p.lookahead().next().expect(token.If)
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

	var lastSpan span.Spanner = thenBlock

	var elseBlock *Block
	if p.lookahead().peek().is(token.Else) {
		p.bump() // consume 'else'

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

		lastSpan = elseBlock
	}

	return ifStmt(
		condition, thenBlock, elseBlock,
		span.Merge(ifTok, lastSpan),
	), nil
}
