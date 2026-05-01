package parser

import "github.com/caiquetorres/lumi/internal/token"

type BreakStmt struct{}

func breakStmt() *BreakStmt {
	return &BreakStmt{}
}

func (p *Parser) parseBreak() (*BreakStmt, error) {
	_, err := p.lookahead().next().expect(token.Break)
	if err != nil {
		return nil, err
	}

	return breakStmt(), nil
}
