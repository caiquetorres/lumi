package lexer

import (
	"errors"
	"io"
	"strings"
)

func (l *Lexer) takeUntil(predicate func(rune) bool) (string, error) {
	var result strings.Builder

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

		_, _ = result.WriteRune(r)
	}

	return result.String(), nil
}

func (l *Lexer) takeWhile(predicate func(rune) bool) (string, error) {
	var result strings.Builder

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

		_, _ = result.WriteRune(r)
	}

	return result.String(), nil
}
