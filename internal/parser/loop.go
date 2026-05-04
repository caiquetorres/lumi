package parser

import "github.com/caiquetorres/lumi/internal/token"

type Loop struct {
	Body *Block
}

func (p *Parser) parseLoop() (*Loop, error) {
	if _, err := p.lookahead().next().expect(token.Loop); err != nil {
		return nil, err
	}

	body, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &Loop{Body: body}, nil
}
