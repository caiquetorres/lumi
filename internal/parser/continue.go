package parser

import "github.com/caiquetorres/lumi/internal/token"

type ContinueStmt struct{}

func continueStmt() *ContinueStmt {
	return &ContinueStmt{}
}

func (p *Parser) parseContinue() (*ContinueStmt, error) {
	_, err := p.lookahead().next().expect(token.Continue)
	if err != nil {
		return nil, err
	}

	return continueStmt(), nil
}
