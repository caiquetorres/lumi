package lexer

import "unicode"

func (l *Lexer) skipWhitespace() error {
	return l.bumpWhile(func(r rune) bool {
		return unicode.IsSpace(r)
	})
}
