package parser

import "github.com/caiquetorres/lumi/internal/token"

func (p *Parser) parseStmt() (Stmt, error) {
	var (
		expr Stmt
		err  error
	)

	switch {
	case p.lookahead().peek().is(token.Let):
		expr, err = p.parseVarDecl()
	case p.lookahead().peek().is(token.Return):
		expr, err = p.parseReturn()
	case p.lookahead().peek().is(token.OpenBrace):
		return p.parseBlock()
	case p.lookahead().peek().is(token.Break):
		expr, err = p.parseBreak()
	default:
		// REVIEW: Can we do better than just trying to parse an expression and returning an error if it fails? Maybe we can check for some other tokens that would indicate that it's not an expression, such as a keyword or a semicolon.

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
