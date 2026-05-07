package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
)

type semanticError struct {
	errs    []error
	isFatal bool
}

func (e *semanticError) Error() string {
	panic("unimplemented")
}

func Analyze(ast *parser.Ast) error {
	return nil
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}
