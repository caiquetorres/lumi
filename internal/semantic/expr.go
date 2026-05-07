package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Expr any

func exprN(e parser.Expr) Expr {
	if e == nil {
		return nil
	}
	switch n := e.(type) {
	case *parser.LiteralExpr:
		return literalExpr(n)
	case *parser.BinaryExpr:
		return binaryExpr(n)
	case *parser.IdentifierExpr:
		return identifierExpr(n)
	case *parser.CallExpr:
		return callExpr(n)
	default:
		panic("unreachable")
	}
}
