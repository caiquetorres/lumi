package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type BreakStmt struct{}

func (t *TypeChecker) analyzeBreakStmt(_ *parser.Break) *BreakStmt {
	return &BreakStmt{}
}
