package lexer

import "unicode"

func (l *Lexer) skipWhitespace() error {
	err := l.bumpWhile(func(r rune) bool {
		return unicode.IsSpace(r)
	})
	l.resetLexeme()
	return err
}
