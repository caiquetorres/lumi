package parser

import (
	"fmt"
	"slices"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

type lookahead struct {
	l   *lexer.Lexer
	buf *circularBuffer
}

func newLookahead(l *lexer.Lexer) *lookahead {
	return &lookahead{
		l:   l,
		buf: newCircularBuffer(2),
	}
}

func (l *lookahead) fill(n uint32) {
	for l.buf.len < n {
		t, err := l.l.Next()
		l.buf.push(&tokenResult{tok: t, err: err})
	}
}

func (l *lookahead) peek() *tokenResult {
	l.fill(1)
	return l.buf.peek(0).(*tokenResult)
}

// func (l *lookahead) peek2() *tokenResult {
// 	l.fill(2)
// 	return l.buf.peek(1).(*tokenResult)
// }

func (l *lookahead) next() *tokenResult {
	l.fill(1)
	return l.buf.pop().(*tokenResult)
}

type tokenResult struct {
	tok token.Token
	err error
}

func (t *tokenResult) get() (token.Token, error) {
	if t.err != nil {
		return token.Token{}, t.err
	}

	return t.tok, nil
}

func (t *tokenResult) is(k token.Kind) bool {
	if t.err != nil {
		return false
	}

	return t.tok.Kind() == k
}

func (t *tokenResult) isOneOf(ks ...token.Kind) bool {
	if t.err != nil {
		return false
	}

	return slices.Contains(ks, t.tok.Kind())
}

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
