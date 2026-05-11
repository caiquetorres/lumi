package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type IdentifierExpr struct {
	Name token.Token

	typedExpr *TypedExpr
}

func identifierExpr(name token.Token, typedExpr *TypedExpr) *IdentifierExpr {
	return &IdentifierExpr{
		typedExpr: typedExpr,
		Name:      name,
	}
}

var _ Expr = (*IdentifierExpr)(nil)

func (i *IdentifierExpr) Type() *TypedExpr {
	return i.typedExpr
}

func (t *TypeChecker) analyzeIdentifierExpr(ie *parser.IdentifierExpr) *IdentifierExpr {
	name := t.lex.Lexeme(ie.Name)
	idExpr := identifierExpr(ie.Name, anyExpr())

	k, exists := t.symTable.Lookup(name)
	if exists {
		idExpr.typedExpr = newTypedExprKindOnly(k)
	}

	return idExpr
}
