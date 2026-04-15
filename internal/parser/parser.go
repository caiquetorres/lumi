package parser

import (
	"fmt"
	"io"
	"slices"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

func New(r io.Reader, l *lexer.Lexer) *Parser {
	if l == nil {
		l = lexer.New(r)
	}
	return &Parser{l: l}
}

type Parser struct {
	l *lexer.Lexer
}

// peek returns the peek token without consuming it. It returns an error
// if there is an error in the lexer.
func (p *Parser) peek() (token.Token, error) {
	return p.l.Peek()
}

// next returns the next token and consumes it. It returns an error
// if there is an error in the lexer.
func (p *Parser) next() (token.Token, error) {
	return p.l.Next()
}

// bump consumes the next token without returning it. It is useful for
// skipping tokens that we don't care about, such as semicolons or commas.
//
// It is unsafe because it ignores any errors from the lexer, so it should
// only be used when we are sure that the next token is of the expected
// kind.
func (p *Parser) bump() {
	_, _ = p.next()
}

// maybeNext checks if the next token is of one of the given kinds, and
// if it is, it consumes it. It is useful for optional tokens, such as
// commas between parameters or semicolons at the end of statements.
func (p *Parser) maybeNext(ks ...token.Kind) {
	if p.peekIsOneOf(ks...) {
		p.bump()
	}
}

// peekIs checks if the next token peekIs of the given kind. It returns
// false if there peekIs an error or if the next token peekIs not of the
// given kind.
func (p *Parser) peekIs(k token.Kind) bool {
	tok, err := p.peek()
	if err != nil {
		return false
	}

	return tok.Kind() == k
}

// peekIsOneOf checks if the next token is of one of the given kinds. It
// returns false if there is an error or if the next token is not of one
// of the given kinds.
func (p *Parser) peekIsOneOf(ks ...token.Kind) bool {
	tok, err := p.peek()
	if err != nil {
		return false
	}

	return slices.Contains(ks, tok.Kind())
}

// expect checks if the next token is of the given kind, and if it is, it
// consumes it and returns it. If the next token is not of the given kind,
// it returns an error.
func (p *Parser) expect(k token.Kind) (token.Token, error) {
	tok, err := p.next()
	if err != nil {
		return token.Token{}, err
	}

	if tok.Kind() != k {
		return token.Token{}, fmt.Errorf("expected token of kind %s, got %s: %w",
			k.String(), tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return tok, nil
}

// expectPeek checks if the next token is of the given kind, and if it is,
// it returns it without consuming it. If the next token is not of the
// given kind, it returns an error.
func (p *Parser) expectPeek(k token.Kind) (token.Token, error) {
	tok, err := p.peek()
	if err != nil {
		return token.Token{}, err
	}

	if tok.Kind() != k {
		return token.Token{}, fmt.Errorf("expected token of kind %s, got %s: %w",
			k.String(), tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return tok, nil
}

// expectSequence checks if the next tokens are of the given kinds, and if
// they are, it consumes them and returns them. If the next tokens are not
// of the given kinds, it returns an error.
func (p *Parser) expectOneOf(ks ...token.Kind) (token.Token, error) {
	tok, err := p.l.Next()
	if err != nil {
		return token.Token{}, err
	}

	if !slices.Contains(ks, tok.Kind()) {
		return token.Token{}, fmt.Errorf("expected token of kind one of %v, got %s: %w",
			ks, tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return tok, nil
}

// expectOneOfPeek checks if the next token is of one of the given kinds, and
// if it is, it returns it without consuming it. If the next token is not of
// one of the given kinds, it returns an error.
func (p *Parser) expectOneOfPeek(ks ...token.Kind) (token.Token, error) {
	tok, err := p.peek()
	if err != nil {
		return token.Token{}, err
	}

	if !slices.Contains(ks, tok.Kind()) {
		return token.Token{}, fmt.Errorf("expected token of kind one of %v, got %s: %w",
			ks, tok.Kind().String(), ErrUnexpectedToken,
		)
	}

	return tok, nil
}

func (p *Parser) expectSequence(ks ...token.Kind) ([]token.Token, error) {
	toks := make([]token.Token, len(ks))
	for i, k := range ks {
		tok, err := p.l.Next()
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

func (p *Parser) err() error {
	_, err := p.peek()
	return err
}

func (p *Parser) expectEndOfLine() error {
	// TODO: add \n for end of line
	_, err := p.expectOneOf(token.Semicolon)
	return err
}
