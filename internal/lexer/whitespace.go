package lexer

import "unicode"

func (l *Lexer) skipWhitespace() error {
	err := l.bumpWhile(func(r rune) bool {
		return unicode.IsSpace(r)
	})
	l.currLexeme = l.currLexeme[:0]
	return err
}
