package lexer

import "unicode"

func (l *Lexer) skipWhitespace() error {
	err := l.bumpWhile(func(r rune) bool {
		if r == '\n' {
			return false
		}
		return unicode.IsSpace(r)
	})
	l.resetLexeme()
	return err
}
