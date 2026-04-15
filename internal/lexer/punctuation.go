package lexer

import (
	"errors"
	"io"

	"github.com/caiquetorres/lumi/internal/token"
)

func (l *Lexer) isPunctuation() bool {
	r, err := l.peekRune()
	if err != nil {
		return false
	}

	_, ok := punctuations[r]
	return ok
}

var punctuations = map[rune]token.Kind{
	'(': token.OpenParen,
	')': token.CloseParen,
	'{': token.OpenBrace,
	'}': token.CloseBrace,
	';': token.Semicolon,
	',': token.Comma,
	'=': token.Equals,
}

func (l *Lexer) readPunctuation() (token.Token, error) {
	r, err := l.nextRune()
	if errors.Is(err, io.EOF) {
		// This should never happen since isPunctuation is called before
		// readPunctuation, but we return an EOF token just in case.
		return l.newToken(token.EOF), nil
	}

	if err != nil {
		return token.Token{}, err
	}

	kind, ok := punctuations[r]
	if !ok {
		// This should never happen since isPunctuation is called before
		// readPunctuation, but we return a Bad token just in case.
		kind = token.Bad
	}

	return l.newToken(kind), nil
}
