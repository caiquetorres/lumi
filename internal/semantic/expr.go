package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Expr any

func (a *Analyzer) analyzeExpr(e parser.Expr) Expr {
	if e == nil {
		return nil
	}
	switch n := e.(type) {
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
