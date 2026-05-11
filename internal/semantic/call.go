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

func (t *TypeChecker) analyzeCallExpr(ce *parser.CallExpr) *CallExpr {
	args := make([]Expr, len(ce.Args))
	for i, arg := range ce.Args {
		args[i] = t.analyzeExpr(arg)
	}

	calle := t.analyzeExpr(ce.Callee)

	typedExpr := anyExpr()
	if calle.Type().IsFunction() {
		calleType := calle.Type().AsFunction()
		typedExpr = newTypedExprKindOnly(calleType.Return)
	}

	return &CallExpr{
		typedExpr: typedExpr,
		Callee:    calle,
		Args:      args,
	}
}
