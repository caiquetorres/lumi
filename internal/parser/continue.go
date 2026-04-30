package parser

import "github.com/caiquetorres/lumi/internal/token"

type Continue struct{}

func (p *Parser) parseContinue() (*Continue, error) {
	_, err := p.lookahead().next().expect(token.Continue)
	if err != nil {
		return nil, err
	}

	return &Continue{}, nil
}
