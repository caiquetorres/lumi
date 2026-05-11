package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
)

type Kind any

type (
	Any      struct{}
	Int      struct{}
	String   struct{}
	Bool     struct{}
	Function struct {
		Params []Kind
		Return Kind
	}
)

func (a *TypeChecker) parseType(t *parser.Type) Kind {
	if t == nil {
		return nil
	}

	name := a.lex.Lexeme(t.Name)

	switch name {
	case "int":
		return Int{}
	case "string":
		return String{}
	case "bool":
		return Bool{}
	default:
		panic("unknown type")
	}
}
