package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type BreakStmt struct{}

func (a *TypeChecker) analyzeBreakStmt(_ *parser.BreakStmt) *BreakStmt {
	return &BreakStmt{}
}
