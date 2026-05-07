package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ContinueStmt struct{}

func continueStmt(_ *parser.ContinueStmt) *ContinueStmt {
	return &ContinueStmt{}
}
