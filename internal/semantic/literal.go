package semantic

import (
	"strconv"

	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type LiteralExpr struct {
	typedExpr *TypedExpr

	Kind  parser.LiteralKind
	Value token.Token
}

func (l *LiteralExpr) Type() *TypedExpr {
	return l.typedExpr
}

var _ Expr = (*LiteralExpr)(nil)

func (t *TypeChecker) analyzeLiteralExpr(lit *parser.LiteralExpr) *LiteralExpr {
	var (
		kind Kind
		val  any
		err  error
	)

	switch lit.Kind {
	case parser.LiteralInt:
		kind = Int{}

		text := t.lex.Lexeme(lit.Value)
		val, err = strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

	case parser.LiteralString:
		kind = String{}

	case parser.LiteralFalse:
		kind = Bool{}
		val = false

	case parser.LiteralTrue:
		kind = Bool{}
		val = true

	default:
		panic("unreachable")
	}

	return &LiteralExpr{
		typedExpr: newTypedExpr(kind, val),
		Kind:      lit.Kind,
		Value:     lit.Value,
	}
}
