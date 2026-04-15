package parser

import "github.com/caiquetorres/lumi/internal/token"

func (p *Parser) parseStmt() (Stmt, error) {
	var (
		expr Stmt
		err  error
	)

	switch {
	case p.peekIs(token.Let):
		expr, err = p.parseVarDecl()
	default:
		expr, err = p.parseExpr()
	}

	if err != nil {
		return nil, err
	}

	if err := p.expectEndOfLine(); err != nil {
		return nil, err
	}

	return expr, nil
}
