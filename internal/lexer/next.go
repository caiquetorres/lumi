package lexer

import (
	"github.com/caiquetorres/lumi/internal/token"
)

// TODO: create a bad token error

func (l *Lexer) next() (token.Token, error) {
	if err := l.skipWhitespace(); err != nil {
		return token.Token{}, err
	}

	switch {
	case l.isAtEOF():
		return l.newToken(token.EOF), nil
	case l.isPunctuation():
		return l.readPunctuation()
	case l.isNumber():
		return l.readNumber()
	case l.isString():
		return l.readString()
	case l.isKeywordOrIdentifier():
		return l.readKeywordOrIdentifier()
	default:
		return l.newToken(token.Bad), nil
	}
}
