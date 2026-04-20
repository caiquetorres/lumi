package parser

import "github.com/caiquetorres/lumi/internal/token"

type BlockExpr struct {
	Stmts []Stmt
}

func (p *Parser) parseBlock() (*BlockExpr, error) {
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

	return &BlockExpr{Stmts: stms}, nil
}
