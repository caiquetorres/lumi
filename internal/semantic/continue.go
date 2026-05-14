package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ContinueStmt struct{}

func (t *TypeChecker) analyzeContinueStmt(_ *parser.Continue) *ContinueStmt {
	return &ContinueStmt{}
}
