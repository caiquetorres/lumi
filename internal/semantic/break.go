package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Break struct{}

func (t *TypeChecker) analyzeBreakStmt(_ *parser.Break) *Break {
	return &Break{}
}
