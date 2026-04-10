package lexer_test

import (
	"strings"
	"testing"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	t.Run("next", func(t *testing.T) {
		l := lexer.New(strings.NewReader("foo bar"))
		tok, err := l.Next()

		assert.NoError(t, err)
		assert.Equal(t, token.Identifier, tok.Kind())
		assert.Equal(t, 0, tok.Span().Start())
		assert.Equal(t, 3, tok.Span().End())
		assert.Equal(t, 3, tok.Span().Len())
	})

	t.Run("peek", func(t *testing.T) {
		l := lexer.New(strings.NewReader("foo bar"))
		tok, err := l.Peek()

		assert.NoError(t, err)
		assert.Equal(t, token.Identifier, tok.Kind())
		assert.Equal(t, 0, tok.Span().Start())
		assert.Equal(t, 3, tok.Span().End())
		assert.Equal(t, 3, tok.Span().Len())

		tok2, err := l.Next()

		assert.NoError(t, err)
		assert.Equal(t, token.Identifier, tok2.Kind())
		assert.Equal(t, 0, tok2.Span().Start())
		assert.Equal(t, 3, tok2.Span().End())
		assert.Equal(t, 3, tok2.Span().Len())
	})

	t.Run("handles multibyte identifier runes", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{name: "multibyte identifier 1", input: "fóo"},
			{name: "multibyte identifier 2", input: "bår"},
			{name: "emoji", input: "😀"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				tok, err := l.Next()

				assert.NoError(t, err, "expected no error when lexing input with multibyte runes")
				assert.Equal(t, token.Identifier, tok.Kind(), "expected token kind to be Identifier")
				assert.Equal(t, 0, tok.Span().Start(), "expected token span to start at 0")
				assert.Equal(t, len([]rune(tt.input)), tok.Span().Len(), "expected token span length to match input length")
			})
		}
	})

	t.Run("punctuation", func(t *testing.T) {
		l := lexer.New(strings.NewReader("(){}"))
		kinds, err := readAllTokenKinds(l)
		expectedKinds := []token.Kind{
			token.OpenParen,
			token.CloseParen,
			token.OpenBrace,
			token.CloseBrace,
			token.EOF,
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedKinds, kinds)
	})

	t.Run("keywords", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			wantKind token.Kind
		}{
			{name: "fun", input: "fun", wantKind: token.Fun},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				kinds, err := readAllTokenKinds(l)
				expectedKinds := []token.Kind{tt.wantKind, token.EOF}

				assert.NoError(t, err)
				assert.Equal(t, expectedKinds, kinds)
			})
		}
	})

	t.Run("sequence of tokens", func(t *testing.T) {
		t.Run("empty input", func(t *testing.T) {
			l := lexer.New(strings.NewReader(""))

			kinds, err := readAllTokenKinds(l)
			expectedKinds := []token.Kind{token.EOF}

			assert.NoError(t, err)
			assert.Equal(t, expectedKinds, kinds)
		})

		t.Run("simple input", func(t *testing.T) {
			l := lexer.New(strings.NewReader("foo bar"))
			kinds, err := readAllTokenKinds(l)
			expectedKinds := []token.Kind{
				token.Identifier,
				token.Identifier,
				token.EOF,
			}

			assert.NoError(t, err)
			assert.Equal(t, expectedKinds, kinds)
		})

		t.Run("simple statement", func(t *testing.T) {
			l := lexer.New(strings.NewReader("fun foo() {}"))

			kinds, err := readAllTokenKinds(l)
			expectedKinds := []token.Kind{
				token.Fun,
				token.Identifier,
				token.OpenParen,
				token.CloseParen,
				token.OpenBrace,
				token.CloseBrace,
				token.EOF,
			}

			assert.NoError(t, err)
			assert.Equal(t, expectedKinds, kinds)
		})
	})
}

func readAllTokenKinds(l *lexer.Lexer) ([]token.Kind, error) {
	var kinds []token.Kind

	for {
		tok, err := l.Next()
		if err != nil {
			return nil, err
		}

		kinds = append(kinds, tok.Kind())

		if tok.Kind() == token.EOF {
			break
		}
	}

	return kinds, nil
}
