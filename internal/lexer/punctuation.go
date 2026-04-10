package lexer

import (
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
}

func (l *Lexer) readPunctuation() token.Token {
	r, err := l.nextRune()
	if err != nil {
		return token.Token{}
	}

	kind, ok := punctuations[r]
	if !ok {
		// This should never happen since isPunctuation is called before
		// readPunctuation, but we return a Bad token just in case.
		kind = token.Bad
	}

	return l.newToken(kind)
}
