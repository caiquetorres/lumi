package parser

import "github.com/caiquetorres/lumi/internal/token"

type Type struct {
	Name token.Token
}

func (p *Parser) isType() bool {
	return p.peek().is(token.Identifier)
}

func (p *Parser) parseType() (*Type, error) {
	tok, err := p.next().expect(token.Identifier)
	if err != nil {
		return nil, err
	}

	return &Type{
		Name: tok,
	}, nil
}
