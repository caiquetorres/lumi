package lexer

import (
	"io"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestBump(t *testing.T) {
	t.Run("consumes one rune", func(t *testing.T) {
		l := newRaw(strings.NewReader("abc"))

		l.bump()

		r, err := l.peekRune()
		assert.NoError(t, err)
		assert.Equal(t, 'b', r)
	})

	t.Run("at eof is a no-op", func(t *testing.T) {
		l := newRaw(strings.NewReader(""))

		l.bump()

		_, err := l.peekRune()
		assert.ErrorIs(t, err, io.EOF)
	})
}

func TestBumpUntil(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		predicate     func(rune) bool
		wantNextRune  rune
		wantNextIsEOF bool
	}{
		{
			name:         "stops before delimiter",
			input:        "abc:def",
			predicate:    func(r rune) bool { return r == ':' },
			wantNextRune: ':',
		},
		{
			name:         "predicate true on first rune",
			input:        ":abc",
			predicate:    func(r rune) bool { return r == ':' },
			wantNextRune: ':',
		},
		{
			name:          "consumes until eof when predicate never matches",
			input:         "abc",
			predicate:     func(r rune) bool { return r == ':' },
			wantNextIsEOF: true,
		},
		{
			name:         "handles multibyte runes",
			input:        "aç9",
			predicate:    func(r rune) bool { return unicode.IsDigit(r) },
			wantNextRune: '9',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newRaw(strings.NewReader(tt.input))

			l.bumpUntil(tt.predicate)

			r, err := l.peekRune()
			if tt.wantNextIsEOF {
				assert.ErrorIs(t, err, io.EOF)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantNextRune, r)
		})
	}
}

func TestBumpWhile(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		predicate     func(rune) bool
		wantNextRune  rune
		wantNextIsEOF bool
	}{
		{
			name:         "consumes while predicate is true",
			input:        "abc:def",
			predicate:    func(r rune) bool { return unicode.IsLetter(r) },
			wantNextRune: ':',
		},
		{
			name:         "predicate false on first rune",
			input:        ":abc",
			predicate:    func(r rune) bool { return unicode.IsLetter(r) },
			wantNextRune: ':',
		},
		{
			name:          "consumes until eof when predicate always true",
			input:         "abc",
			predicate:     func(r rune) bool { return true },
			wantNextIsEOF: true,
		},
		{
			name:         "handles multibyte runes",
			input:        "aç9",
			predicate:    func(r rune) bool { return unicode.IsLetter(r) },
			wantNextRune: '9',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newRaw(strings.NewReader(tt.input))

			l.bumpWhile(tt.predicate)

			r, err := l.peekRune()
			if tt.wantNextIsEOF {
				assert.ErrorIs(t, err, io.EOF)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantNextRune, r)
		})
	}
}
