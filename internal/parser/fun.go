package parser

import (
	"io"

	"github.com/caiquetorres/lumi/internal/token"
)

type FunDecl struct {
	Identifier token.Token
	Body       []Expression
}

func (p *Parser) debugFuncDel(f *FunDecl, w io.Writer) error {
	if _, err := w.Write([]byte("fun")); err != nil {
		return err
	}

	w.Write([]byte(" "))

	if _, err := w.Write(p.l.Lexeme(f.Identifier)); err != nil {
		return err
	}

	w.Write([]byte(" "))

	if _, err := w.Write([]byte("()")); err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseFunDecl() (*FunDecl, error) {
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

	return &FunDecl{
		Identifier: toks[1],
		Body:       body,
	}, nil
}

func (p *Parser) parseFunDeclBody() ([]Expression, error) {
	body := make([]Expression, 0)

	_, err := p.expectSequence(token.OpenBrace, token.CloseBrace)
	if err != nil {
		return nil, err
	}

	return body, nil
}

var _ TopLevelStmt = (*FunDecl)(nil)
