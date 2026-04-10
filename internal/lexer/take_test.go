package lexer

import (
	"io"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestTakeUntil(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		predicate     func(rune) bool
		wantText      string
		wantStart     int
		wantEnd       int
		wantNextRune  rune
		wantNextIsEOF bool
	}{
		{
			name:         "stops before delimiter",
			input:        "abc:def",
			predicate:    func(r rune) bool { return r == ':' },
			wantText:     "abc",
			wantStart:    0,
			wantEnd:      3,
			wantNextRune: ':',
		},
		{
			name:         "predicate true on first rune",
			input:        ":abc",
			predicate:    func(r rune) bool { return r == ':' },
			wantText:     "",
			wantStart:    0,
			wantEnd:      0,
			wantNextRune: ':',
		},
		{
			name:          "reads until eof when predicate never matches",
			input:         "abc",
			predicate:     func(r rune) bool { return r == ':' },
			wantText:      "abc",
			wantStart:     0,
			wantEnd:       3,
			wantNextIsEOF: true,
		},
		{
			name:         "counts runes not bytes",
			input:        "aç9",
			predicate:    func(r rune) bool { return unicode.IsDigit(r) },
			wantText:     "aç",
			wantStart:    0,
			wantEnd:      2,
			wantNextRune: '9',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newRaw(strings.NewReader(tt.input))

			gotText, _ := l.takeUntil(tt.predicate)

			assert.Equal(t, tt.wantText, gotText)
			assert.Equal(t, tt.wantStart, l.start)
			assert.Equal(t, tt.wantEnd, l.end)

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

func TestTakeWhile(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		predicate     func(rune) bool
		wantText      string
		wantStart     int
		wantEnd       int
		wantNextRune  rune
		wantNextIsEOF bool
	}{
		{
			name:         "reads while predicate is true",
			input:        "abc:de",
			predicate:    func(r rune) bool { return unicode.IsLetter(r) },
			wantText:     "abc",
			wantStart:    0,
			wantEnd:      3,
			wantNextRune: ':',
		},
		{
			name:         "predicate false on first rune",
			input:        ":abc",
			predicate:    func(r rune) bool { return unicode.IsLetter(r) },
			wantText:     "",
			wantStart:    0,
			wantEnd:      0,
			wantNextRune: ':',
		},
		{
			name:          "reads full input until eof when predicate always true",
			input:         "abc",
			predicate:     func(r rune) bool { return true },
			wantText:      "abc",
			wantStart:     0,
			wantEnd:       3,
			wantNextIsEOF: true,
		},
		{
			name:         "counts runes not bytes",
			input:        "aç9",
			predicate:    func(r rune) bool { return unicode.IsLetter(r) },
			wantText:     "aç",
			wantStart:    0,
			wantEnd:      2,
			wantNextRune: '9',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newRaw(strings.NewReader(tt.input))

			gotText, _ := l.takeWhile(tt.predicate)

			assert.Equal(t, tt.wantText, gotText)
			assert.Equal(t, tt.wantStart, l.start)
			assert.Equal(t, tt.wantEnd, l.end)

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
