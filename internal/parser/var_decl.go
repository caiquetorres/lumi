package parser

import "github.com/caiquetorres/lumi/internal/token"

type VarDecl struct {
	Identifier token.Token
	Expr       Expr
}

func (p *Parser) parseVarDecl() (*VarDecl, error) {
	// let <identifier> = <expr>

	toks, err := p.expectSequence(token.Let, token.Identifier, token.Equals)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return &VarDecl{
		Identifier: toks[1],
		Expr:       expr,
	}, nil
}
