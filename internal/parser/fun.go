package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type FunDecl struct {
	Identifier token.Token
	Body       []Expr
}

func (p *Parser) parseFunDecl() (*FunDecl, error) {
	// func <identifier>() { <body> }
	// func <identifier>()

	toks, err := p.expectSequence(token.Fun, token.Identifier,
		token.OpenParen, token.CloseParen)
	if err != nil {
		return nil, err
	}

	var body []Expr
	if p.is(token.OpenBrace) {
		// The function body is optional, so we only parse it if we see an
		// opening brace.

		body, err = p.parseFunDeclBody()
		if err != nil {
			return nil, err
		}
	}

	return &FunDecl{
		Identifier: toks[1],
		Body:       body,
	}, nil
}

func (p *Parser) parseFunDeclBody() ([]Expr, error) {
	body := make([]Expr, 0)

	_, err := p.expect(token.OpenBrace)
	if err != nil {
		return nil, err
	}

	for !p.is(token.CloseBrace) {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if _, err := p.expect(token.Semicolon); err != nil {
			return nil, err
		}

		body = append(body, expr)
	}

	_, err = p.expect(token.CloseBrace)
	if err != nil {
		return nil, err
	}

	return body, nil
}

var _ TopLevelStmt = (*FunDecl)(nil)
