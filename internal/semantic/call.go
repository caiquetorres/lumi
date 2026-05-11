package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func (a *Analyzer) analyzeCallExpr(ce *parser.CallExpr) *CallExpr {
	args := make([]Expr, len(ce.Args))
	for i, arg := range ce.Args {
		args[i] = a.analyzeExpr(arg)
	}
	return &CallExpr{
		Callee: a.analyzeExpr(ce.Callee),
		Args:   args,
	}
}
