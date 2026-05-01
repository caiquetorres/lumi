package parser

import "github.com/caiquetorres/lumi/internal/token"

type Block struct {
	Stmts []Stmt
}

func block(stmts []Stmt) *Block {
	return &Block{
		Stmts: stmts,
	}
}

var _ Stmt = (*Block)(nil)

func (p *Parser) parseBlock() (*Block, error) {
	_, err := p.lookahead().next().expect(token.OpenBrace)
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

		stms = append(stms, stmt)
	}

	if p.lookahead().peek().is(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	p.bump() // consume the '}'

	return block(stms), nil
}
