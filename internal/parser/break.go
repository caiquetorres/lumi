package parser

import "github.com/caiquetorres/lumi/internal/token"

type Break struct {
	Expr Expr
}

func (p *Parser) parseBreak() (*Break, error) {
	_, err := p.lookahead().next().expect(token.Break)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return &Break{
		Expr: expr,
	}, nil
}
