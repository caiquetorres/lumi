package parser

import (
	"fmt"
	"io"
	"slices"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

func Parse(r io.Reader) (*Ast, error) {
	p := new(r)
	return p.parseAst()
}

type parser struct {
	l *lexer.Lexer
}

func new(r io.Reader) *parser {
	l := lexer.New(r)
	return &parser{l: l}
}

func (p *parser) is(k token.Kind) bool {
	tok, err := p.peek()
	if err != nil {
		return false
	}

	return tok.Kind() == k
}

func (p *parser) isOneOf(ks ...token.Kind) bool {
	tok, err := p.peek()
	if err != nil {
		return false
	}

	return slices.Contains(ks, tok.Kind())
}

func (p *parser) peek() (token.Token, error) {
	return p.l.Peek()
}

func (p *parser) expect(k token.Kind) (token.Token, error) {
	tok, err := p.l.Next()
	if err != nil {
		return token.Token{}, err
	}

	if tok.Kind() != k {
		return token.Token{}, fmt.Errorf("expected token of kind %s, got %s", k.String(), tok.Kind().String())
	}

	return tok, nil
}

func (p *parser) expectOneOf(ks ...token.Kind) (token.Token, error) {
	tok, err := p.l.Next()
	if err != nil {
		return token.Token{}, err
	}

	if !slices.Contains(ks, tok.Kind()) {
		return token.Token{}, fmt.Errorf("expected token of kind one of %v, got %s", ks, tok.Kind().String())
	}

	return tok, nil
}

func (p *parser) expectSequence(ks ...token.Kind) ([]token.Token, error) {
	toks := make([]token.Token, len(ks))
	for i, k := range ks {
		tok, err := p.expect(k)
		if err != nil {
			return nil, fmt.Errorf("expected token of kind %s at position %d: %w", k.String(), i, err)
		}

		toks[i] = tok
	}

	return toks, nil
}
