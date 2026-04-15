package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type FunDecl struct {
	Identifier token.Token
	Params     []Param
	Body       []Stmt
}

type Param struct {
	Name token.Token
	Type Type
}

func (p *Parser) parseFunDecl() (*FunDecl, error) {
	// func <identifier>() { <body> }
	// func <identifier>()

	toks, err := p.expectSequence(token.Fun, token.Identifier, token.OpenParen)
	if err != nil {
		return nil, err
	}

	var params []Param
	for !p.peekIs(token.CloseParen) {
		param, err := p.parseParam()
		if err != nil {
			return nil, err
		}

		params = append(params, *param)

		_, err = p.expectOneOfPeek(token.Comma, token.CloseParen, token.EOF)
		if err != nil {
			return nil, err
		}

		p.maybeNext(token.Comma)
	}

	if p.peekIs(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	p.bump() // consume the ')'

	var body []Stmt
	if p.peekIs(token.OpenBrace) {
		// The function body is optional, so we only parse it if we see an
		// opening brace.

		body, err = p.parseFunDeclBody()
		if err != nil {
			return nil, err
		}
	}

	return &FunDecl{
		Identifier: toks[1],
		Params:     params,
		Body:       body,
	}, nil
}

func (p *Parser) parseParam() (*Param, error) {
	tok, err := p.expect(token.Identifier)
	if err != nil {
		return nil, err
	}

	var ty *Type
	if p.isType() {
		ty, err = p.parseType()
		if err != nil {
			return nil, err
		}
	}

	return &Param{
		Name: tok,
		Type: *ty,
	}, nil
}

func (p *Parser) parseFunDeclBody() ([]Stmt, error) {
	if _, err := p.expect(token.OpenBrace); err != nil {
		return nil, err
	}

	body := make([]Stmt, 0)

	for !p.peekIsOneOf(token.CloseBrace, token.EOF) {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	if p.peekIs(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	p.bump() // consume the '}'

	return body, nil
}

var _ TopLevelStmt = (*FunDecl)(nil)
