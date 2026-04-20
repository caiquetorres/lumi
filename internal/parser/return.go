package parser

import "github.com/caiquetorres/lumi/internal/token"

type Return struct {
	Expr Expr
}

func (p *Parser) parseReturn() (*Return, error) {
	_, err := p.lookahead().next().expect(token.Return)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return &Return{
		Expr: expr,
	}, nil
}
