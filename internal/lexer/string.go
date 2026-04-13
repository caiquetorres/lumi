package lexer

import "github.com/caiquetorres/lumi/internal/token"

func (l *Lexer) isString() bool {
	r, err := l.peekRune()
	if err != nil {
		return false
	}

	return r == '"'
}

func (l *Lexer) readString() (token.Token, error) {
	if _, err := l.nextRune(); err != nil {
		return token.Token{}, err
	}

	_, err := l.takeWhile(func(r rune) bool {
		return r != '"'
	})
	if err != nil {
		return token.Token{}, err
	}

	if _, err := l.nextRune(); err != nil {
		return token.Token{}, err
	}

	return l.newToken(token.String), nil
}
