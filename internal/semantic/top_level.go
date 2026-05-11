package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type TopLevelStmt any

type Ast struct {
	Statements []TopLevelStmt
}

func (a *TypeChecker) analyzeTopLevelStmt(s parser.TopLevelStmt) TopLevelStmt {
	switch n := s.(type) {
	case *parser.FunDecl:
		return a.analyzeFunDecl(n)
	case *parser.VarDecl:
		return a.analyzeVarDecl(n)
	default:
		panic("unreachable")
	}
}
