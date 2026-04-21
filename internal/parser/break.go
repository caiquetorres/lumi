package parser

import "github.com/caiquetorres/lumi/internal/token"

type Break struct{}

func (p *Parser) parseBreak() (*Break, error) {
	_, err := p.lookahead().next().expect(token.Break)
	if err != nil {
		return nil, err
	}

	return &Break{}, nil
}
