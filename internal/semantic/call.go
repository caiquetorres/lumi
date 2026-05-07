package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func callExpr(ce *parser.CallExpr) *CallExpr {
	args := make([]Expr, len(ce.Args))
	for i, a := range ce.Args {
		args[i] = exprN(a)
	}
	return &CallExpr{
		Callee: exprN(ce.Callee),
		Args:   args,
	}
}
