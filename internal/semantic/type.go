package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type Type struct {
	Name token.Token
}

func (a *TypeChecker) analyzeType(ty *parser.Type) *Type {
	if ty == nil {
		return nil
	}
	return &Type{
		Name: ty.Name,
	}
}
