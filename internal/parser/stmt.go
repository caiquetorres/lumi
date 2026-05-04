package parser

import "github.com/caiquetorres/lumi/internal/token"

type Stmt any

func (p *Parser) parseStmt() (Stmt, error) {
	var (
		stmt Stmt
		err  error
	)

	switch {
	case p.lookahead().peek().is(token.Let):
		stmt, err = p.parseVarDecl()
	case p.lookahead().peek().is(token.If):
		stmt, err = p.parseIf()
	case p.lookahead().peek().is(token.While):
		stmt, err = p.parseWhile()
	case p.lookahead().peek().is(token.Return):
		stmt, err = p.parseReturn()
	case p.lookahead().peek().is(token.OpenBrace):
		return p.parseBlock()
	case p.lookahead().peek().is(token.Break):
		stmt, err = p.parseBreak()
	case p.lookahead().peek().is(token.Continue):
		stmt, err = p.parseContinue()
	case p.lookahead().peek().is(token.For):
		stmt, err = p.parseFor()
	default:
		// REVIEW: Can we do better than just trying to parse an expression and returning an error if it fails? Maybe we can check for some other tokens that would indicate that it's not an expression, such as a keyword or a semicolon.

		stmt, err = p.parseExpr()
	}

	if err != nil {
		return nil, err
	}

	return stmt, nil
}
