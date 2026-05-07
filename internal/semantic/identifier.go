package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type IdentifierExpr struct {
	Name token.Token
}

func identifierExpr(ie *parser.IdentifierExpr) *IdentifierExpr {
	return &IdentifierExpr{
		Name: ie.Name,
	}
}
