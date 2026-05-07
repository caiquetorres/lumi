package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type TopLevelStmt any

type Ast struct {
	Statements []TopLevelStmt
}

func topLevelStmt(s parser.TopLevelStmt) TopLevelStmt {
	switch n := s.(type) {
	case *parser.FunDecl:
		return funDecl(n)
	case *parser.VarDecl:
		return varDecl(n)
	default:
		panic("unreachable")
	}
}
