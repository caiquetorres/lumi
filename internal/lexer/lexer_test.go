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
		assert.Equal(t, uint32(0), tok.Span().Start())
		assert.Equal(t, uint32(3), tok.Span().End())
		assert.Equal(t, uint32(3), tok.Span().Len())
	})

	t.Run("peek", func(t *testing.T) {
		l := lexer.New(strings.NewReader("foo bar"))
		tok, err := l.Peek()

		assert.NoError(t, err)
		assert.Equal(t, token.Identifier, tok.Kind())
		assert.Equal(t, uint32(0), tok.Span().Start())
		assert.Equal(t, uint32(3), tok.Span().End())
		assert.Equal(t, uint32(3), tok.Span().Len())

		tok2, err := l.Next()

		assert.NoError(t, err)
		assert.Equal(t, token.Identifier, tok2.Kind())
		assert.Equal(t, uint32(0), tok2.Span().Start())
		assert.Equal(t, uint32(3), tok2.Span().End())
		assert.Equal(t, uint32(3), tok2.Span().Len())
	})

	t.Run("identifiers", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{name: "leading underscore", input: "_foo"},
			{name: "underscore in the middle", input: "foo_bar"},
			{name: "only underscore", input: "_"},
			{name: "digit at the end", input: "foo1"},
			{name: "digit in the middle", input: "x2y"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				kinds, err := readAllTokenKinds(l)

				assert.NoError(t, err)
				assert.Equal(t, []token.Kind{token.Identifier, token.EOF}, kinds)
			})
		}
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
				assert.Equal(t, uint32(0), tok.Span().Start(), "expected token span to start at 0")
				assert.Equal(t, uint32(len([]rune(tt.input))), tok.Span().Len(), "expected token span length to match input length")
			})
		}
	})

	t.Run("bad", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{name: "at sign", input: "@"},
			{name: "hash", input: "#"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				tok, err := l.Next()

				assert.NoError(t, err)
				assert.Equal(t, token.Bad, tok.Kind())
			})
		}
	})

	t.Run("punctuation", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			wantKind token.Kind
		}{
			{name: "open paren", input: "(", wantKind: token.OpenParen},
			{name: "close paren", input: ")", wantKind: token.CloseParen},
			{name: "open brace", input: "{", wantKind: token.OpenBrace},
			{name: "close brace", input: "}", wantKind: token.CloseBrace},
			{name: "semicolon", input: ";", wantKind: token.Semicolon},
			{name: "comma", input: ",", wantKind: token.Comma},
			{name: "newline", input: "\n", wantKind: token.NewLine},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				kinds, err := readAllTokenKinds(l)

				assert.NoError(t, err)
				assert.Equal(t, []token.Kind{tt.wantKind, token.EOF}, kinds)
			})
		}

		t.Run("sequence of punctuations", func(t *testing.T) {
			l := lexer.New(strings.NewReader("(){};,"))
			kinds, err := readAllTokenKinds(l)
			expectedKinds := []token.Kind{
				token.OpenParen,
				token.CloseParen,
				token.OpenBrace,
				token.CloseBrace,
				token.Semicolon,
				token.Comma,
				token.EOF,
			}

			assert.NoError(t, err)
			assert.Equal(t, expectedKinds, kinds)
		})
	})

	t.Run("keywords", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			wantKind token.Kind
		}{
			{name: "fun", input: "fun", wantKind: token.Fun},
			{name: "let", input: "let", wantKind: token.Let},
			{name: "return", input: "return", wantKind: token.Return},
			{name: "true", input: "true", wantKind: token.True},
			{name: "false", input: "false", wantKind: token.False},
			{name: "if", input: "if", wantKind: token.If},
			{name: "else", input: "else", wantKind: token.Else},
			{name: "loop", input: "loop", wantKind: token.Loop},
			{name: "while", input: "while", wantKind: token.While},
			{name: "break", input: "break", wantKind: token.Break},
			{name: "continue", input: "continue", wantKind: token.Continue},
			{name: "for", input: "for", wantKind: token.For},
			{name: "in", input: "in", wantKind: token.In},
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

		t.Run("keyword as identifier prefix is an identifier", func(t *testing.T) {
			tests := []struct {
				name  string
				input string
			}{
				{name: "fun prefix", input: "function"},
				{name: "let prefix", input: "letter"},
				{name: "if prefix", input: "iffy"},
				{name: "for prefix", input: "forEach"},
				{name: "in prefix", input: "index"},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					l := lexer.New(strings.NewReader(tt.input))
					kinds, err := readAllTokenKinds(l)
					expectedKinds := []token.Kind{token.Identifier, token.EOF}

					assert.NoError(t, err)
					assert.Equal(t, expectedKinds, kinds)
				})
			}
		})
	})

	t.Run("numbers", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{name: "single digit", input: "0"},
			{name: "multiple digits", input: "123"},
			{name: "leading zeros", input: "007"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				kinds, err := readAllTokenKinds(l)

				assert.NoError(t, err)
				assert.Equal(t, []token.Kind{token.Int, token.EOF}, kinds)
			})
		}

		t.Run("sequence of numbers", func(t *testing.T) {
			l := lexer.New(strings.NewReader("1 2 3"))
			kinds, err := readAllTokenKinds(l)

			assert.NoError(t, err)
			assert.Equal(t, []token.Kind{token.Int, token.Int, token.Int, token.EOF}, kinds)
		})
	})

	t.Run("strings", func(t *testing.T) {
		tests := []struct {
			name     string
			input    string
			wantKind token.Kind
		}{
			{name: "simple string", input: `"hello"`, wantKind: token.String},
			{name: "empty string", input: `""`, wantKind: token.String},
			{name: "string with spaces", input: `"hello world"`, wantKind: token.String},
			{name: "string with numbers", input: `"abc123"`, wantKind: token.String},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				kinds, err := readAllTokenKinds(l)

				assert.NoError(t, err)
				assert.Equal(t, []token.Kind{tt.wantKind, token.EOF}, kinds)
			})
		}

		t.Run("unclosed string returns error", func(t *testing.T) {
			l := lexer.New(strings.NewReader(`"unclosed`))
			_, err := readAllTokenKinds(l)

			assert.Error(t, err)
		})
	})

	t.Run("operators", func(t *testing.T) {
		t.Run("simple", func(t *testing.T) {
			tests := []struct {
				name     string
				input    string
				wantKind token.Kind
			}{
				{name: "plus", input: "+", wantKind: token.Plus},
				{name: "minus", input: "-", wantKind: token.Minus},
				{name: "star", input: "*", wantKind: token.Star},
				{name: "slash", input: "/", wantKind: token.Slash},
				{name: "equal", input: "=", wantKind: token.Equal},
				{name: "bang", input: "!", wantKind: token.Bang},
				{name: "less", input: "<", wantKind: token.Less},
				{name: "greater", input: ">", wantKind: token.Greater},
				{name: "dot", input: ".", wantKind: token.Dot},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					l := lexer.New(strings.NewReader(tt.input))
					kinds, err := readAllTokenKinds(l)

					assert.NoError(t, err)
					assert.Equal(t, []token.Kind{tt.wantKind, token.EOF}, kinds)
				})
			}
		})

		t.Run("compound", func(t *testing.T) {
			tests := []struct {
				name     string
				input    string
				wantKind token.Kind
			}{
				{name: "plus equal", input: "+=", wantKind: token.PlusEqual},
				{name: "minus equal", input: "-=", wantKind: token.MinusEqual},
				{name: "star equal", input: "*=", wantKind: token.StarEqual},
				{name: "slash equal", input: "/=", wantKind: token.SlashEqual},
				{name: "equal equal", input: "==", wantKind: token.EqualEqual},
				{name: "bang equal", input: "!=", wantKind: token.BangEqual},
				{name: "less equal", input: "<=", wantKind: token.LessEqual},
				{name: "greater equal", input: ">=", wantKind: token.GreaterEqual},
				{name: "dot dot", input: "..", wantKind: token.DotDot},
				{name: "dot dot equal", input: "..=", wantKind: token.DotDotEqual},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					l := lexer.New(strings.NewReader(tt.input))
					kinds, err := readAllTokenKinds(l)

					assert.NoError(t, err)
					assert.Equal(t, []token.Kind{tt.wantKind, token.EOF}, kinds)
				})
			}
		})
	})

	t.Run("whitespace", func(t *testing.T) {
		t.Run("spaces between tokens are skipped", func(t *testing.T) {
			l := lexer.New(strings.NewReader("foo   bar"))
			kinds, err := readAllTokenKinds(l)

			assert.NoError(t, err)
			assert.Equal(t, []token.Kind{token.Identifier, token.Identifier, token.EOF}, kinds)
		})

		t.Run("tabs between tokens are skipped", func(t *testing.T) {
			l := lexer.New(strings.NewReader("foo\t\tbar"))
			kinds, err := readAllTokenKinds(l)

			assert.NoError(t, err)
			assert.Equal(t, []token.Kind{token.Identifier, token.Identifier, token.EOF}, kinds)
		})

		t.Run("newline is not skipped", func(t *testing.T) {
			l := lexer.New(strings.NewReader("foo\nbar"))
			kinds, err := readAllTokenKinds(l)

			assert.NoError(t, err)
			assert.Equal(t, []token.Kind{token.Identifier, token.NewLine, token.Identifier, token.EOF}, kinds)
		})

		t.Run("leading and trailing spaces are skipped", func(t *testing.T) {
			l := lexer.New(strings.NewReader("  foo  "))
			kinds, err := readAllTokenKinds(l)

			assert.NoError(t, err)
			assert.Equal(t, []token.Kind{token.Identifier, token.EOF}, kinds)
		})
	})

	t.Run("spans", func(t *testing.T) {
		t.Run("number span", func(t *testing.T) {
			l := lexer.New(strings.NewReader("42"))
			tok, err := l.Next()

			assert.NoError(t, err)
			assert.Equal(t, token.Int, tok.Kind())
			assert.Equal(t, uint32(0), tok.Span().Start())
			assert.Equal(t, uint32(2), tok.Span().End())
			assert.Equal(t, uint32(2), tok.Span().Len())
		})

		t.Run("string span", func(t *testing.T) {
			l := lexer.New(strings.NewReader(`"hi"`))
			tok, err := l.Next()

			assert.NoError(t, err)
			assert.Equal(t, token.String, tok.Kind())
			assert.Equal(t, uint32(0), tok.Span().Start())
			assert.Equal(t, uint32(4), tok.Span().End())
			assert.Equal(t, uint32(4), tok.Span().Len())
		})

		t.Run("compound operator span", func(t *testing.T) {
			l := lexer.New(strings.NewReader("+="))
			tok, err := l.Next()

			assert.NoError(t, err)
			assert.Equal(t, token.PlusEqual, tok.Kind())
			assert.Equal(t, uint32(0), tok.Span().Start())
			assert.Equal(t, uint32(2), tok.Span().End())
			assert.Equal(t, uint32(2), tok.Span().Len())
		})

		t.Run("span includes preceding whitespace", func(t *testing.T) {
			l := lexer.New(strings.NewReader("a + b"))

			a, _ := l.Next()  // "a"
			op, _ := l.Next() // " +" (whitespace + operator)
			b, _ := l.Next()  // " b"

			assert.Equal(t, uint32(0), a.Span().Start())
			assert.Equal(t, uint32(1), a.Span().End())

			assert.Equal(t, uint32(1), op.Span().Start())
			assert.Equal(t, uint32(3), op.Span().End())

			assert.Equal(t, uint32(3), b.Span().Start())
			assert.Equal(t, uint32(5), b.Span().End())
		})
	})

	t.Run("lexeme", func(t *testing.T) {
		tests := []struct {
			name       string
			input      string
			wantLexeme string
		}{
			{name: "identifier", input: "foo", wantLexeme: "foo"},
			{name: "number", input: "123", wantLexeme: "123"},
			{name: "string", input: `"hello"`, wantLexeme: `"hello"`},
			{name: "keyword", input: "let", wantLexeme: "let"},
			{name: "operator", input: "+", wantLexeme: "+"},
			{name: "compound operator", input: "+=", wantLexeme: "+="},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(strings.NewReader(tt.input))
				tok, err := l.Next()

				assert.NoError(t, err)
				assert.Equal(t, tt.wantLexeme, l.Lexeme(tok))
			})
		}

		t.Run("same string interns to same lexeme", func(t *testing.T) {
			l := lexer.New(strings.NewReader("foo foo"))

			tok1, _ := l.Next()
			tok2, _ := l.Next()

			assert.Equal(t, l.Lexeme(tok1), l.Lexeme(tok2))
			assert.Equal(t, tok1.SymbolID(), tok2.SymbolID())
		})
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

		t.Run("arithmetic expression", func(t *testing.T) {
			l := lexer.New(strings.NewReader("a + b"))
			kinds, err := readAllTokenKinds(l)

			assert.NoError(t, err)
			assert.Equal(t, []token.Kind{token.Identifier, token.Plus, token.Identifier, token.EOF}, kinds)
		})

		t.Run("variable declaration", func(t *testing.T) {
			l := lexer.New(strings.NewReader("let x = 42"))
			kinds, err := readAllTokenKinds(l)

			assert.NoError(t, err)
			assert.Equal(t, []token.Kind{token.Let, token.Identifier, token.Equal, token.Int, token.EOF}, kinds)
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
