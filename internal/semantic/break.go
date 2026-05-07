package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type BreakStmt struct{}

func breakStmt(_ *parser.BreakStmt) *BreakStmt {
	return &BreakStmt{}
}
