package parser

import (
	"fmt"
	"slices"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

const lookaheadBufferSize = 2

// lookahead is a helper struct that allows us to peek at the next token
// without consuming it.
type lookahead struct {
	lex *lexer.Lexer
	buf *circularBuffer
}

func newLookahead(l *lexer.Lexer) *lookahead {
	return &lookahead{
		lex: l,
		buf: newCircularBuffer(lookaheadBufferSize),
	}
}

// fill fills the lookahead buffer with tokens until it has at least n tokens.
func (l *lookahead) fill(n uint32) {
	for l.buf.len < n {
		t, err := l.lex.Next()
		l.buf.pushBack(&tokenResult{tok: t, err: err})
	}
}

// peek returns the next token without consuming it.
func (l *lookahead) peek() *tokenResult {
	l.fill(1)
	return l.buf.at(0).(*tokenResult)
}

// peek2 returns the second next token without consuming it.
func (l *lookahead) peek2() *tokenResult {
	l.fill(2)
	return l.buf.at(1).(*tokenResult)
}

// next consumes the next token and returns it.
func (l *lookahead) next() *tokenResult {
	l.fill(1)
	return l.buf.popFront().(*tokenResult)
}

// tokenResult is a helper struct that wraps a token and an error, and
// provides some helper methods for checking the token kind and expecting
// certain kinds of tokens.
type tokenResult struct {
	tok token.Token
	err error
}

// get returns the token if there was no error, or an error if there was
// an error.
func (t *tokenResult) get() (token.Token, error) {
	if t.err != nil {
		return token.Token{}, t.err
	}

	return t.tok, nil
}

// is checks if the token is of the given kind. It returns false if there
// was an error, or if the token is of a different kind.
func (t *tokenResult) is(k token.Kind) bool {
	if t.err != nil {
		return false
	}

	return t.tok.Kind() == k
}

// isOneOf checks if the token is of one of the given kinds. It returns
// false if there was an error, or if the token is of a different kind.
func (t *tokenResult) isOneOf(ks ...token.Kind) bool {
	if t.err != nil {
		return false
	}

	return slices.Contains(ks, t.tok.Kind())
}

// expect checks if the token is of the given kind, and returns it if it
// is. Otherwise, it returns an error. It also returns an error if there
// was an error in the token result.
func (t *tokenResult) expect(k token.Kind) (token.Token, error) {
	if t.err != nil {
		return token.Token{}, t.err
	}

	if !t.is(k) {
		return token.Token{}, fmt.Errorf("expected token of kind %s, got %s: %w",
			k.String(), t.tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return t.tok, nil
}

// expectOneOf checks if the token is of one of the given kinds, and
// returns it if it is. Otherwise, it returns an error. It also returns an
// error if there was an error in the token result.
func (t *tokenResult) expectOneOf(ks ...token.Kind) (token.Token, error) {
	if t.err != nil {
		return token.Token{}, t.err
	}

	if !t.isOneOf(ks...) {
		return token.Token{}, fmt.Errorf("expected token of one of kinds %v, got %s: %w",
			ks, t.tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return t.tok, nil
}
