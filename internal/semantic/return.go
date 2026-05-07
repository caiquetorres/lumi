package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ReturnStmt struct {
	Expr Expr
}

func returnStmt(rs *parser.ReturnStmt) *ReturnStmt {
	return &ReturnStmt{
		Expr: exprN(rs.Expr),
	}
}
