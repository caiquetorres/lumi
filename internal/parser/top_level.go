package parser

import (
	"io"

	"github.com/caiquetorres/lumi/internal/token"
)

type TopLevelStmt interface{}

func (p *Parser) debugTopLevelStmt(stmt TopLevelStmt, w io.Writer) error {
	switch s := stmt.(type) {
	case *FunDecl:
		return p.debugFuncDel(s, w)
	default:
		panic("unreachable")
	}
}

// parseTopLevelStmt parses a top-level statement, which can be: a function
// declaration, a variable declaration, a package declaration, or an import
// statement.
func (p *Parser) parseTopLevelStmt() (TopLevelStmt, error) {
	switch {
	case p.is(token.Fun):
		return p.parseFunDecl()
	default:
		if err := p.err(); err != nil {
			return nil, err
		}

		return p.expectOneOf(token.Fun)
	}
}
