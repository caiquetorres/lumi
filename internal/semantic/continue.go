package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ContinueStmt struct{}

func (t *TypeChecker) analyzeContinueStmt(cs *parser.ContinueStmt) *ContinueStmt {
	return &ContinueStmt{}
}
