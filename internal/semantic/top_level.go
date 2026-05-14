package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type TopLevelStmt any

type Ast struct {
	Statements []TopLevelStmt
}

func (t *TypeChecker) analyzeTopLevelStmt(s parser.TopLevelStmt) TopLevelStmt {
	switch n := s.(type) {
	case *parser.FunDecl:
		return t.analyzeFunDecl(n)
	case *parser.Let:
		return t.analyzeVarDecl(n)
	default:
		panic("unreachable")
	}
}
