package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Expr interface {
	expr() // marker method
}

func (p *Parser) parseExpr() (Expr, error) {
	return p.parseBinaryExpr(-1)
}

func (p *Parser) parseUnit() (Expr, error) {
	switch {
	case p.lookahead().peek().is(token.OpenParen):
		p.bump() // consume '('

		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		_, err = p.lookahead().next().expect(token.CloseParen)
		if err != nil {
			return nil, err
		}

		return expr, nil
	case p.isLiteral():
		return p.parseLiteral()
	case p.lookahead().peek().is(token.Identifier):
		return p.parseIdentifier()
	default:
		_, err := p.
			lookahead().
			peek().
			expectOneOf(
				token.String, token.Int, token.True, token.False,
				token.OpenParen,
			)

		return nil, err
	}
}
