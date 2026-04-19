package parser

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	return p
}

type Parser struct {
	l *lexer.Lexer

	t token.Token
	e error
}

// bump consumes the next token without returning it. It is useful for
// skipping tokens that we don't care about, such as semicolons or commas.
//
// It is unsafe because it ignores any errors from the lexer, so it should
// only be used when we are sure that the next token is of the expected
// kind.
func (p *Parser) bump() {
	_ = p.next()
}

// maybeNext checks if the next token is of one of the given kinds, and
// if it is, it consumes it. It is useful for optional tokens, such as
// commas between parameters or semicolons at the end of statements.
func (p *Parser) maybeNext(ks ...token.Kind) {
	if p.peek().isOneOf(ks...) {
		p.bump()
	}
}

func (p *Parser) expectSequence(ks ...token.Kind) ([]token.Token, error) {
	toks := make([]token.Token, len(ks))
	for i, k := range ks {
		tok, err := p.next().get()
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

func (p *Parser) expectEndOfLine() error {
	// TODO: add \n for end of line
	_, err := p.next().expectOneOf(token.Semicolon)
	return err
}
