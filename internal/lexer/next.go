package lexer

import (
	"github.com/caiquetorres/lumi/internal/token"
)

func (l *Lexer) next() (token.Token, error) {
	l.skipWhitespace()

	switch {
	case l.isAtEOF():
		return l.newToken(token.EOF), nil
	case l.isPunctuation():
		return l.readPunctuation(), nil
	case l.isKeywordOrIdentifier():
		return l.readKeywordOrIdentifier(), nil
	default:
		return l.newToken(token.Bad), nil
	}
}
