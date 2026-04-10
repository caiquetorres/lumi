package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Fun struct {
	Identifier token.Token
	Body       []Expression
}

func (p *parser) parseFunDecl() (*Fun, error) {
	// func <identifier>() { <body> }
	// func <identifier>()

	toks, err := p.expectSequence(token.Fun, token.Identifier,
		token.OpenParen, token.CloseParen)
	if err != nil {
		return nil, err
	}

	body := make([]Expression, 0)
	if p.is(token.OpenBrace) {
		// The function body is optional, so we only parse it if we see an
		// opening brace.

		body, err = p.parseFunDeclBody()
		if err != nil {
			return nil, err
		}
	}

	return &Fun{
		Identifier: toks[1],
		Body:       body,
	}, nil
}

func (p *parser) parseFunDeclBody() ([]Expression, error) {
	body := make([]Expression, 0)

	if p.is(token.OpenBrace) {
		_, err := p.expectSequence(token.OpenBrace, token.CloseBrace)
		if err != nil {
			return nil, err
		}
	}

	return body, nil
}

var _ TopLevelStmt = (*Fun)(nil)
