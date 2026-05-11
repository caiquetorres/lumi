package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type CallExpr struct {
	typedExpr *TypedExpr

	Callee Expr
	Args   []Expr
}

var _ Expr = (*CallExpr)(nil)

func (c *CallExpr) Type() *TypedExpr {
	return c.typedExpr
}

func (a *TypeChecker) analyzeCallExpr(ce *parser.CallExpr) *CallExpr {
	args := make([]Expr, len(ce.Args))
	for i, arg := range ce.Args {
		args[i] = a.analyzeExpr(arg)
	}
	return &CallExpr{
		typedExpr: anyExpr(),
		Callee:    a.analyzeExpr(ce.Callee),
		Args:      args,
	}
}
