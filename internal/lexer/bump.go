package lexer

import (
	"errors"
	"io"
)

func (l *Lexer) bump() {
	_, _ = l.nextRune()
}

func (l *Lexer) bumpUntil(fn func(rune) bool) error {
	for {
		r, err := l.peekRune()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		if fn(r) {
			break
		}

		l.bump()
	}

	return nil
}

func (l *Lexer) bumpWhile(fn func(rune) bool) error {
	for {
		r, err := l.peekRune()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		if !fn(r) {
			break
		}

		l.bump()
	}

	return nil
}
