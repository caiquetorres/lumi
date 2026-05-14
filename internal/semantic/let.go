package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type Let struct {
	Bindings []Binding
}

func (t *TypeChecker) analyzeLet(le *parser.Let) *Let {
	assignments := make([]Binding, len(le.Bindings))
	for i, as := range le.Bindings {
		assignments[i] = t.analyzeBinding(as)
	}
	return &Let{
		Bindings: assignments,
	}
}

type Binding struct {
	Identifier token.Token
	Expr       Expr
}

func (t *TypeChecker) analyzeBinding(bi parser.Binding) Binding {
	name := t.lex.Lexeme(bi.Identifier)
	t.symTable.Define(name, bi.Expr)

	return Binding{
		Identifier: bi.Identifier,
		Expr:       t.analyzeExpr(bi.Expr),
	}
}
