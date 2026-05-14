package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Continue struct{}

func (t *TypeChecker) analyzeContinueStmt(_ *parser.Continue) *Continue {
	return &Continue{}
}
