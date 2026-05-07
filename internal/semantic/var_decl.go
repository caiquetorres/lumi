package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type VarDecl struct {
	Assignments []Assignment
}

func varDecl(vd *parser.VarDecl) *VarDecl {
	assignments := make([]Assignment, len(vd.Assignments))
	for i, a := range vd.Assignments {
		assignments[i] = assignment(a)
	}
	return &VarDecl{
		Assignments: assignments,
	}
}

type Assignment struct {
	Identifier token.Token
	Expr       Expr
}

func assignment(a parser.Assignment) Assignment {
	return Assignment{
		Identifier: a.Identifier,
		Expr:       exprN(a.Expr),
	}
}
