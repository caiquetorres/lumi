package parser

import (
	"fmt"
	"io"
	"slices"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

func New(r io.Reader, l *lexer.Lexer) *Parser {
	if l == nil {
		l = lexer.New(r)
	}
	return &Parser{l: l}
}

type Parser struct {
	l *lexer.Lexer
}

func (p *Parser) is(k token.Kind) bool {
	tok, err := p.peek()
	if err != nil {
		return false
	}

	return tok.Kind() == k
}

func (p *Parser) isOneOf(ks ...token.Kind) bool {
	tok, err := p.peek()
	if err != nil {
		return false
	}

	return slices.Contains(ks, tok.Kind())
}

func (p *Parser) err() error {
	_, err := p.peek()
	return err
}

func (p *Parser) peek() (token.Token, error) {
	return p.l.Peek()
}

func (p *Parser) next() (token.Token, error) {
	return p.l.Next()
}

func (p *Parser) expect(k token.Kind) (token.Token, error) {
	tok, err := p.l.Next()
	if err != nil {
		return token.Token{}, err
	}

	if tok.Kind() != k {
		return token.Token{}, fmt.Errorf("expected token of kind %s, got %s: %w",
			k.String(), tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return tok, nil
}

func (p *Parser) expectOneOf(ks ...token.Kind) (token.Token, error) {
	tok, err := p.l.Next()
	if err != nil {
		return token.Token{}, err
	}

	if !slices.Contains(ks, tok.Kind()) {
		return token.Token{}, fmt.Errorf("expected token of kind one of %v, got %s: %w",
			ks, tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return tok, nil
}

func (p *Parser) expectSequence(ks ...token.Kind) ([]token.Token, error) {
	toks := make([]token.Token, len(ks))
	for i, k := range ks {
		tok, err := p.l.Next()
		if err != nil {
			return nil, err
		}

		if tok.Kind() != k {
			return nil, fmt.Errorf("expected token of kind %s at position %d: %w",
				k.String(), i, ErrUnexpectedToken)
		}

		toks[i] = tok
	}

	return toks, nil
}

func (p *Parser) expectEndOfLine() error {
	// TODO: add \n for end of line
	_, err := p.expectOneOf(token.Semicolon)
	return err
}
