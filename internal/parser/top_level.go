package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type TopLevelStmt any

// parseTopLevelStmt parses a top-level statement, which can be: a function
// declaration, a variable declaration, a package declaration, or an import
// statement.
func (p *Parser) parseTopLevelStmt() (TopLevelStmt, error) {
	switch {
	case p.peekIs(token.Fun):
		return p.parseFunDecl()
	default:
		return p.expectOneOf(token.Fun)
	}
}
