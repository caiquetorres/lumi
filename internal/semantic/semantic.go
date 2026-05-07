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

func Analyze(ast *parser.Ast) (*Ast, error) {
	return astN(ast), nil
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func astN(a *parser.Ast) *Ast {
	stmts := make([]TopLevelStmt, len(a.Statements))
	for i, s := range a.Statements {
		stmts[i] = topLevelStmt(s)
	}
	return &Ast{
		Statements: stmts,
	}
}
