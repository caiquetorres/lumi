package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type VarDecl struct {
	Assignments []Assignment
}

func (t *TypeChecker) analyzeVarDecl(vd *parser.VarDecl) *VarDecl {
	assignments := make([]Assignment, len(vd.Assignments))
	for i, as := range vd.Assignments {
		assignments[i] = t.analyzeAssignment(as)
	}
	return &VarDecl{
		Assignments: assignments,
	}
}

type Assignment struct {
	Identifier token.Token
	Expr       Expr
}

func (t *TypeChecker) analyzeAssignment(as parser.Assignment) Assignment {
	name := t.lex.Lexeme(as.Identifier)
	t.symTable.Define(name, as.Expr)

	return Assignment{
		Identifier: as.Identifier,
		Expr:       t.analyzeExpr(as.Expr),
	}
}
