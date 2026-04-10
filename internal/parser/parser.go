package parser

import (
	"io"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

func Parse(r io.Reader) (*Ast, error) {
	_ = new(r)
	return &Ast{}, nil
}

type parser struct{}

func (p *parser) expect(k token.Kind) (token.Token, error) {
	return token.Token{}, nil
}

func (p *parser) expectSequence(ks ...token.Kind) ([]token.Token, error) {
	return nil, nil
}

func new(r io.Reader) *parser {
	_ = lexer.New(r)
	return &parser{}
}
