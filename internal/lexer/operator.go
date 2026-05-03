package lexer

import (
	"github.com/caiquetorres/lumi/internal/token"
)

var operators = map[rune]struct{}{
	'+': {},
	'-': {},
	'*': {},
	'/': {},
	'=': {},
	'!': {},
	'<': {},
	'>': {},
	'.': {},
}

func (l *Lexer) isOperator() bool {
	r, err := l.peekRune()
	if err != nil {
		return false
	}

	_, ok := operators[r]
	return ok
}

func (l *Lexer) readOperator() (token.Token, error) {
	r, err := l.nextRune()
	if err != nil {
		return token.Token{}, err
	}

	switch r {
	case '+':
		return l.readPlusOrPlusEqual()
	case '-':
		return l.readMinusOrMinusEqual()
	case '*':
		return l.readStarOrStarEqual()
	case '/':
		return l.readSlashOrSlashEqual()
	case '=':
		return l.readEqualOrEqualEqual()
	case '!':
		return l.readBangOrBangEqual()
	case '<':
		return l.readLessOrLessEqual()
	case '>':
		return l.readGreaterOrGreaterEqual()
	case '.':
		return l.readDotOrDotDotOrDotDotEqual()
	}

	return token.Token{}, nil
}

func (l *Lexer) readPlusOrPlusEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.PlusEqual), nil
	}

	return l.newToken(token.Plus), nil
}

func (l *Lexer) readMinusOrMinusEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.MinusEqual), nil
	}

	return l.newToken(token.Minus), nil
}

func (l *Lexer) readStarOrStarEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.StarEqual), nil
	}

	return l.newToken(token.Star), nil
}

func (l *Lexer) readSlashOrSlashEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.SlashEqual), nil
	}

	return l.newToken(token.Slash), nil
}

func (l *Lexer) readEqualOrEqualEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.EqualEqual), nil
	}

	return l.newToken(token.Equal), nil
}

func (l *Lexer) readBangOrBangEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.BangEqual), nil
	}

	return l.newToken(token.Bang), nil
}

func (l *Lexer) readGreaterOrGreaterEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.GreaterEqual), nil
	}

	return l.newToken(token.Greater), nil
}

func (l *Lexer) readLessOrLessEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '=' {
		l.bump() // consume the '='
		return l.newToken(token.LessEqual), nil
	}

	return l.newToken(token.Less), nil
}

func (l *Lexer) readDotOrDotDotOrDotDotEqual() (token.Token, error) {
	r, err := l.peekRune()
	if err != nil {
		return token.Token{}, err
	}

	if r == '.' {
		l.bump() // consume the second '.'

		r, err := l.peekRune()
		if err != nil {
			return token.Token{}, err
		}

		if r == '=' {
			l.bump() // consume the '='
			return l.newToken(token.DotDotEqual), nil
		}

		return l.newToken(token.DotDot), nil
	}

	return l.newToken(token.Dot), nil
}
