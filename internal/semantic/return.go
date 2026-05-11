package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ReturnStmt struct {
	Expr Expr
}

func (a *Analyzer) analyzeReturnStmt(rs *parser.ReturnStmt) *ReturnStmt {
	return &ReturnStmt{
		Expr: a.analyzeExpr(rs.Expr),
	}
}
