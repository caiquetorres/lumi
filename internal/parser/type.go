package parser

import "github.com/caiquetorres/lumi/internal/token"

type Type struct {
	Name token.Token
}

func typeN(name token.Token) *Type {
	return &Type{
		Name: name,
	}
}

func (p *Parser) isType() bool {
	return p.lookahead().peek().is(token.Identifier)
}

func (p *Parser) parseType() (*Type, error) {
	tok, err := p.lookahead().next().expect(token.Identifier)
	if err != nil {
		return nil, err
	}

	return typeN(tok), nil
}
