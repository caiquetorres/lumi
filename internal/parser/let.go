package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type Let struct {
	Bindings []Binding

	span span.Span
}

func letStmt(bindings []Binding, span span.Spanner) *Let {
	return &Let{
		Bindings: bindings,
		span:     span.Span(),
	}
}

func (s *Let) Span() span.Span {
	return s.span
}

func (p *Parser) parseLet() (*Let, error) {
	// let <identifier_1> = <expr>, <identifier_2> = <expr>, ...

	letTok, err := p.lookahead().next().expect(token.Let)
	if err != nil {
		return nil, err
	}

	assignments := make([]Binding, 0)

	var lastSpan span.Spanner

	hasNext := true
	for hasNext {
		bi, err := p.parseBinding()
		if err != nil {
			return nil, err
		}

		assignments = append(assignments, *bi)

		lastSpan = bi.Expr

		hasNext = p.lookahead().peek().is(token.Comma)
		if hasNext {
			_, _ = p.lookahead().next().get()
		}
	}

	return letStmt(assignments, span.Merge(letTok, lastSpan)), nil
}

type Binding struct {
	Identifier token.Token
	Expr       Expr
}

func binding(idenfier token.Token, expr Expr) *Binding {
	return &Binding{
		Identifier: idenfier,
		Expr:       expr,
	}
}

func (p *Parser) parseBinding() (*Binding, error) {
	toks, err := p.expectSequence(token.Identifier, token.Equal)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return binding(toks[0], expr), nil
}
