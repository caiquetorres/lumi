package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type VarDecl struct {
	Assignments []Assignment
}

func (a *TypeChecker) analyzeVarDecl(vd *parser.VarDecl) *VarDecl {
	assignments := make([]Assignment, len(vd.Assignments))
	for i, as := range vd.Assignments {
		assignments[i] = a.analyzeAssignment(as)
	}
	return &VarDecl{
		Assignments: assignments,
	}
}

type Assignment struct {
	Identifier token.Token
	Expr       Expr
}

func (a *TypeChecker) analyzeAssignment(as parser.Assignment) Assignment {
	name := a.lex.Lexeme(as.Identifier)
	a.symTable.Define(name, as.Expr)

	return Assignment{
		Identifier: as.Identifier,
		Expr:       a.analyzeExpr(as.Expr),
	}
}
