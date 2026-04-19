package parser

import (
	"fmt"
	"slices"

	"github.com/caiquetorres/lumi/internal/token"
)

type tokenResult struct {
	tok token.Token
	e   error
}

func (p *Parser) peek() *tokenResult {
	p.t, p.e = p.l.Peek()

	return &tokenResult{
		tok: p.t,
		e:   p.e,
	}
}

func (p *Parser) next() *tokenResult {
	p.t, p.e = p.l.Next()

	return &tokenResult{
		tok: p.t,
		e:   p.e,
	}
}

func (t *tokenResult) err() error {
	return t.e
}

func (t *tokenResult) get() (token.Token, error) {
	if t.err() != nil {
		return token.Token{}, t.e
	}

	return t.tok, nil
}

func (t *tokenResult) is(k token.Kind) bool {
	if t.err() != nil {
		return false
	}

	return t.tok.Kind() == k
}

func (t *tokenResult) isOneOf(ks ...token.Kind) bool {
	if t.err() != nil {
		return false
	}

	return slices.Contains(ks, t.tok.Kind())
}

func (t *tokenResult) expect(k token.Kind) (token.Token, error) {
	if t.err() != nil {
		return token.Token{}, t.e
	}

	if !t.is(k) {
		return token.Token{}, fmt.Errorf("expected token of kind %s, got %s: %w",
			k.String(), t.tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return t.tok, nil
}

func (t *tokenResult) expectOneOf(ks ...token.Kind) (token.Token, error) {
	if t.err() != nil {
		return token.Token{}, t.e
	}

	if !t.isOneOf(ks...) {
		return token.Token{}, fmt.Errorf("expected token of one of kinds %v, got %s: %w",
			ks, t.tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return t.tok, nil
}
