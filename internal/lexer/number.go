package lexer

import (
	"unicode"

	"github.com/caiquetorres/lumi/internal/token"
)

func (l *Lexer) isNumber() bool {
	r, err := l.peekRune()
	if err != nil {
		return false
	}

	return unicode.IsDigit(r)
}

func (l *Lexer) readNumber() (token.Token, error) {
	_, err := l.takeWhile(func(r rune) bool {
		return unicode.IsDigit(r)
	})
	if err != nil {
		return token.Token{}, err
	}

	return l.newToken(token.Int), nil
}
