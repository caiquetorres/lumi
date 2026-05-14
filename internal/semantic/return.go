package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ReturnStmt struct {
	Expr Expr
}

func (t *TypeChecker) analyzeReturnStmt(re *parser.Return) *ReturnStmt {
	return &ReturnStmt{
		Expr: t.analyzeExpr(re.Expr),
	}
}
