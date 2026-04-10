package parser

import "github.com/caiquetorres/lumi/internal/token"

type Fun struct {
	Identifier token.Token
	Body       []Expression
}

func (p *parser) parseFun() (*Fun, error) {
	toks, err := p.expectSequence(
		token.Fun,
		token.Identifier,
		token.OpenBrace,
		token.CloseBrace,
	)
	if err != nil {
		return nil, err
	}

	return &Fun{
		Identifier: toks[1],
		Body:       make([]Expression, 0),
	}, nil
}

var _ TopLevelStmt = (*Fun)(nil)
