package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type Block struct {
	Stmts []Stmt

	span span.Span
}

func block(stmts []Stmt, span span.Span) *Block {
	return &Block{
		Stmts: stmts,
		span:  span,
	}
}

func (s *Block) Span() span.Span {
	return s.span
}

var _ Stmt = (*Block)(nil)

func (p *Parser) parseBlock() (*Block, error) {
	openBraceTok, err := p.lookahead().next().expect(token.OpenBrace)
	if err != nil {
		return nil, err
	}

	stms := make([]Stmt, 0)

	for {
		p.skipWhitespace()

		if p.lookahead().peek().isOneOf(token.CloseBrace) {
			break
		}

		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		if err := p.expectEndOfLine(); err != nil {
			return nil, err
		}

		stms = append(stms, stmt)
	}

	if p.lookahead().peek().is(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	closeBraceTok, err := p.lookahead().next().expect(token.CloseBrace)
	if err != nil {
		return nil, err
	}

	return block(stms, span.Merge(openBraceTok, closeBraceTok)), nil
}
