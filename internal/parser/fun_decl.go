package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type FunDecl struct {
	Identifier token.Token
	Params     []Param
	Body       []Stmt
	Return     *Type
}

func funDecl(identifier token.Token, params []Param, body []Stmt, returnType *Type) *FunDecl {
	return &FunDecl{
		Identifier: identifier,
		Params:     params,
		Body:       body,
		Return:     returnType,
	}
}

func (p *Parser) parseFunDecl() (*FunDecl, error) {
	// func <identifier>()
	// func <identifier>() { <body> }
	// func <identifier>() -> <type> { <body> }
	// func <identifier>(param1, param2, ...) { <body> }
	// func <identifier>(param1, param2, ...) -> <type> { <body> }

	toks, err := p.expectSequence(token.Fun, token.Identifier, token.OpenParen)
	if err != nil {
		return nil, err
	}

	identifier := toks[1]

	var params []Param
	for !p.lookahead().peek().isOneOf(token.CloseParen, token.EOF) {
		param, err := p.parseParam()
		if err != nil {
			return nil, err
		}

		params = append(params, *param)

		_, err = p.lookahead().peek().expectOneOf(token.Comma, token.CloseParen)
		if err != nil {
			return nil, err
		}

		if p.lookahead().peek().is(token.Comma) {
			p.bump() // close paren
		}
	}

	if p.lookahead().peek().is(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	p.bump() // consume the ')'

	var ty *Type
	if p.isType() {
		ty, err = p.parseType()
		if err != nil {
			return nil, err
		}
	}

	var body []Stmt
	if p.lookahead().peek().is(token.OpenBrace) {
		// The function body is optional, so we only parse it if we see an
		// opening brace.

		body, err = p.parseFunDeclBody()
		if err != nil {
			return nil, err
		}
	}

	return funDecl(identifier, params, body, ty), nil
}

type Param struct {
	Name token.Token
	Type Type
}

func param(name token.Token, ty Type) *Param {
	return &Param{
		Name: name,
		Type: ty,
	}
}

func (p *Parser) parseParam() (*Param, error) {
	// <identifier>
	// <identifier> <type>

	tok, err := p.lookahead().next().expect(token.Identifier)
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

	return param(tok, *ty), nil
}

func (p *Parser) parseFunDeclBody() ([]Stmt, error) {
	_, err := p.lookahead().next().expect(token.OpenBrace)
	if err != nil {
		return nil, err
	}

	body := make([]Stmt, 0)

	for {
		p.skipWhitespace()

		if p.lookahead().peek().isOneOf(token.CloseBrace, token.EOF) {
			break
		}

		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	if p.lookahead().peek().is(token.EOF) {
		return nil, ErrUnexpectedEOF
	}

	p.bump() // consume the '}'

	return body, nil
}

var _ TopLevelStmt = (*FunDecl)(nil)
