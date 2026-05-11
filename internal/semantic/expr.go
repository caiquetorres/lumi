package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Expr interface {
	Type() *TypedExpr
}

func (t *TypeChecker) analyzeExpr(exp parser.Expr) Expr {
	if exp == nil {
		return nil
	}

	switch n := exp.(type) {
	case *parser.LiteralExpr:
		return t.analyzeLiteralExpr(n)
	case *parser.BinaryExpr:
		return t.analyzeBinaryExpr(n)
	case *parser.IdentifierExpr:
		return t.analyzeIdentifierExpr(n)
	case *parser.CallExpr:
		return t.analyzeCallExpr(n)
	default:
		panic("unreachable")
	}
}
