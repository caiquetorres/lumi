package lexer

import (
	"unicode"

	"github.com/caiquetorres/lumi/internal/token"
)

var keywords = map[string]token.Kind{
	"fun":    token.Fun,
	"let":    token.Let,
	"return": token.Return,
	"break":  token.Break,
	"true":   token.True,
	"false":  token.False,
}

func (l *Lexer) isKeywordOrIdentifier() bool {
	r, err := l.peekRune()
	if err != nil {
		return false
	}

	return unicode.IsLetter(r) || unicode.IsMark(r) ||
		unicode.IsSymbol(r) || r == '_'
}

func (l *Lexer) readKeywordOrIdentifier() (token.Token, error) {
	text, err := l.takeWhile(func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r) ||
			unicode.IsMark(r) || unicode.IsSymbol(r) || r == '_'
	})
	if err != nil {
		return token.Token{}, err
	}

	kind, exists := keywords[text]
	if !exists {
		kind = token.Identifier
	}

	return l.newToken(kind), nil
}
