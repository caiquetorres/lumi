package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type BreakStmt struct{}

func (a *Analyzer) analyzeBreakStmt(bs *parser.BreakStmt) *BreakStmt {
	return &BreakStmt{}
}
