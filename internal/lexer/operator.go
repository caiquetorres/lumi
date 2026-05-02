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
		return l.newToken(token.Plus), nil
	case '-':
		return l.newToken(token.Minus), nil
	case '*':
		return l.newToken(token.Star), nil
	case '/':
		return l.newToken(token.Slash), nil
	case '=':
		return l.readEqualOrEqualEqual()
	case '!':
		return l.readBangOrBangEqual()
	case '<':
		return l.readLessOrLessEqual()
	case '>':
		return l.readGreaterOrGreaterEqual()
	}

	return token.Token{}, nil
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
