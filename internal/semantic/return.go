package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ReturnStmt struct {
	Expr Expr
}

func (t *TypeChecker) analyzeReturnStmt(rs *parser.ReturnStmt) *ReturnStmt {
	return &ReturnStmt{
		Expr: t.analyzeExpr(rs.Expr),
	}
}
