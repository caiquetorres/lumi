package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type FunDecl struct {
	Identifier token.Token
	Body       []Stmt
}

func (p *Parser) parseFunDecl() (*FunDecl, error) {
	// func <identifier>() { <body> }
	// func <identifier>()

	toks, err := p.expectSequence(token.Fun, token.Identifier,
		token.OpenParen, token.CloseParen)
	if err != nil {
		return nil, err
	}

	var body []Stmt
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

func (p *Parser) parseFunDeclBody() ([]Stmt, error) {
	if _, err := p.expect(token.OpenBrace); err != nil {
		return nil, err
	}

	body := make([]Stmt, 0)

	for !p.isOneOf(token.CloseBrace, token.EOF) {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	if p.is(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	if _, err := p.expect(token.CloseBrace); err != nil {
		return nil, err
	}

	return body, nil
}

var _ TopLevelStmt = (*FunDecl)(nil)
