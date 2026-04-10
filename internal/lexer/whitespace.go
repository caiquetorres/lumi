package lexer

import "unicode"

func (l *Lexer) skipWhitespace() {
	l.bumpWhile(func(r rune) bool {
		return unicode.IsSpace(r)
	})
}
