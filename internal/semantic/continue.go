package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ContinueStmt struct{}

func (a *Analyzer) analyzeContinueStmt(cs *parser.ContinueStmt) *ContinueStmt {
	return &ContinueStmt{}
}
