package semantic

import (
	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

type semanticError struct {
	errs    []error
	isFatal bool
}

func (e *semanticError) Error() string {
	panic("unimplemented")
}

func Analyze(ast *parser.Ast, lex *lexer.Lexer) (*Ast, error) {
	stmts := make([]TopLevelStmt, len(ast.Statements))

	analyzer := NewAnalyzer(lex)
	for i, s := range ast.Statements {
		stmts[i] = analyzer.analyzeTopLevelStmt(s)
	}

	return &Ast{
		Statements: stmts,
	}, nil
}
