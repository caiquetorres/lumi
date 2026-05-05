package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type VarDecl struct {
	Assignments []Assignment

	span span.Span
}

type Assignment struct {
	Identifier token.Token
	Expr       Expr
}

func varDeclStmt(assignments []Assignment, span span.Spanner) *VarDecl {
	return &VarDecl{
		Assignments: assignments,
		span:        span.Span(),
	}
}

func (s *VarDecl) Span() span.Span {
	return s.span
}

func (p *Parser) parseVarDecl() (*VarDecl, error) {
	// let <identifier_1> = <expr>, <identifier_2> = <expr>, ...

	letTok, err := p.lookahead().next().expect(token.Let)
	if err != nil {
		return nil, err
	}

	hasNext := true
	assignments := make([]Assignment, 0)

	var lastSpan span.Spanner

	for hasNext {
		toks, err := p.expectSequence(token.Identifier, token.Equal)
		if err != nil {
			return nil, err
		}

		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		lastSpan = expr

		assignments = append(assignments, Assignment{
			Identifier: toks[0],
			Expr:       expr,
		})

		hasNext = p.lookahead().peek().is(token.Comma)
		if hasNext {
			_, _ = p.lookahead().next().get()
		}
	}

	return varDeclStmt(assignments, span.Merge(letTok, lastSpan)), nil
}
