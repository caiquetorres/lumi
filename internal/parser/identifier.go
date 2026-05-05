package parser

import (
	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type IdentifierExpr struct {
	Name token.Token
}

func identifierExpr(name token.Token) *IdentifierExpr {
	return &IdentifierExpr{
		Name: name,
	}
}

func (l *IdentifierExpr) expr() {}

func (l *IdentifierExpr) Span() span.Span {
	return l.Name.Span()
}

var _ Expr = (*IdentifierExpr)(nil)

func (p *Parser) parseIdentifier() (*IdentifierExpr, error) {
	tok, err := p.lookahead().next().expect(token.Identifier)
	if err != nil {
		return nil, err
	}

	return identifierExpr(tok), nil
}
