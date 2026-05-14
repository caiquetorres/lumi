package lexer

import (
	"unicode"

	"github.com/caiquetorres/lumi/internal/token"
)

var keywords = map[string]token.Kind{
	"fun":      token.Fun,
	"let":      token.Let,
	"return":   token.Return,
	"true":     token.True,
	"false":    token.False,
	"if":       token.If,
	"else":     token.Else,
	"loop":     token.Loop,
	"while":    token.While,
	"break":    token.Break,
	"continue": token.Continue,
	"for":      token.For,
	"in":       token.In,
}

// isKeywordOrIdentifier checks if the next token is a keyword or an
// identifier.
func (l *Lexer) isKeywordOrIdentifier() bool {
	r, err := l.peekRune()
	if err != nil {
		return false
	}

	return isIdentifierStart(r)
}

// readKeywordOrIdentifier reads a keyword or an identifier from the input
// and returns the corresponding token. If the text matches a keyword, the
// token kind will be set to the keyword's kind; otherwise, it will be set
// to [token.Identifier].
func (l *Lexer) readKeywordOrIdentifier() (token.Token, error) {
	text, err := l.takeWhile(isIdentifierContinue)
	if err != nil {
		return token.Token{}, err
	}

	kind, exists := keywords[text]
	if !exists {
		kind = token.Identifier
	}

	return l.newToken(kind), nil
}

func isIdentifierStart(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsMark(r) ||
		unicode.IsSymbol(r) || r == '_'
}

func isIdentifierContinue(r rune) bool {
	return isIdentifierStart(r) || unicode.IsNumber(r)
}
