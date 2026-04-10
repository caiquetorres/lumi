package lexer

import (
	"errors"
	"io"
)

func (l *Lexer) takeUntil(predicate func(rune) bool) (string, error) {
	var result []rune

	for {
		r, err := l.peekRune()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return "", err
		}

		if predicate(r) {
			break
		}

		l.bump()

		result = append(result, r)
	}

	return string(result), nil
}

func (l *Lexer) takeWhile(predicate func(rune) bool) (string, error) {
	var result []rune

	for {
		r, err := l.peekRune()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return "", err
		}

		if !predicate(r) {
			break
		}

		l.bump()

		result = append(result, r)
	}

	return string(result), nil
}
