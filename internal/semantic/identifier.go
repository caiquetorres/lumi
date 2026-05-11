package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type IdentifierExpr struct {
	typedExpr *TypedExpr

	Name token.Token
}

var _ Expr = (*IdentifierExpr)(nil)

func (i *IdentifierExpr) Type() *TypedExpr {
	return i.typedExpr
}

func (a *TypeChecker) analyzeIdentifierExpr(ie *parser.IdentifierExpr) *IdentifierExpr {
	return &IdentifierExpr{
		typedExpr: anyExpr(),
		Name:      ie.Name,
	}
}
