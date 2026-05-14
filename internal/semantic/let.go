package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type Let struct {
	Assignments []Assignment
}

func (t *TypeChecker) analyzeVarDecl(vd *parser.Let) *Let {
	assignments := make([]Assignment, len(vd.Bindings))
	for i, as := range vd.Bindings {
		assignments[i] = t.analyzeAssignment(as)
	}
	return &Let{
		Assignments: assignments,
	}
}

type Assignment struct {
	Identifier token.Token
	Expr       Expr
}

func (t *TypeChecker) analyzeAssignment(as parser.Binding) Assignment {
	name := t.lex.Lexeme(as.Identifier)
	t.symTable.Define(name, as.Expr)

	return Assignment{
		Identifier: as.Identifier,
		Expr:       t.analyzeExpr(as.Expr),
	}
}
