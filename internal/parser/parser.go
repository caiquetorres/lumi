package parser

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

type Parser struct {
	la *lookahead
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		la: newLookahead(l),
	}
}

func (p *Parser) lookahead() *lookahead {
	return p.la
}

// bump consumes the next token without returning it. It is useful for
// skipping tokens that we don't care about, such as semicolons or commas.
//
// It is unsafe because it ignores any errors from the lexer, so it should
// only be used when we are sure that the next token is of the expected
// kind.
func (p *Parser) bump() {
	_ = p.lookahead().next()
}

// skipWhitespace consumes any semicolons or newlines until it finds a
// non-whitespace token. It is useful for ignoring optional semicolons at
// the end of statements, or for allowing multiple statements on the same
// line.
func (p *Parser) skipWhitespace() { // REVIEW: Can we rename this?
	for p.lookahead().peek().isOneOf(token.Semicolon, token.NewLine) {
		p.bump()
	}
}

func (p *Parser) expectSequence(ks ...token.Kind) ([]token.Token, error) {
	// TODO: This is a bit of a hack, we should probably have a better way to expect a sequence of tokens, but for now this will do. We should deprecate this function in favor of using expect and expectOneOf.

	toks := make([]token.Token, len(ks))
	for i, k := range ks {
		tok, err := p.lookahead().next().get()
		if err != nil {
			return nil, err
		}

		if tok.Kind() != k {
			return nil, fmt.Errorf("expected token of kind %s at position %d: %w",
				k.String(), i, ErrUnexpectedToken)
		}

		toks[i] = tok
	}

	return toks, nil
}

// expectEndOfLine checks if the next token is a semicolon or a newline,
// and consumes it.
func (p *Parser) expectEndOfLine() error {
	_, err := p.lookahead().next().expectOneOf(
		token.Semicolon, token.NewLine,
	)
	return err
}
