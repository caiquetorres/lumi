package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Expr interface {
	Type() *TypedExpr
}

func (a *TypeChecker) analyzeExpr(exp parser.Expr) Expr {
	if exp == nil {
		return nil
	}

	switch n := exp.(type) {
	case *parser.LiteralExpr:
		return a.analyzeLiteralExpr(n)
	case *parser.BinaryExpr:
		return a.analyzeBinaryExpr(n)
	case *parser.IdentifierExpr:
		return a.analyzeIdentifierExpr(n)
	case *parser.CallExpr:
		return a.analyzeCallExpr(n)
	default:
		panic("unreachable")
	}
}
