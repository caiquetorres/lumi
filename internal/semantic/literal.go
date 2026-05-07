package semantic

import (
	"strconv"

	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type LiteralExpr struct {
	Kind  parser.LiteralKind
	Value token.Token
}

func literalExpr(le *parser.LiteralExpr) *LiteralExpr {
	return &LiteralExpr{
		Kind:  le.Kind,
		Value: le.Value,
	}
}

func (a *Analyzer) AnalyzeLiteral(lit *parser.LiteralExpr) (*AnalyzedExpr, error) {
	var (
		kind Kind
		val  any
		err  error
	)

	switch lit.Kind {
	case parser.LiteralInt:
		kind = Int{}

		text := a.lex.Lexeme(lit.Value)
		val, err = strconv.Atoi(text)
		if err != nil {
			return nil, err
		}

	case parser.LiteralString:
		kind = String{}

	case parser.LiteralFalse, parser.LiteralTrue:
		kind = Bool{}
		text := a.lex.Lexeme(lit.Value)

		val = text == "true"
	default:
		panic("unreachable")
	}

	return &AnalyzedExpr{
		Kind:  kind,
		Value: val,
		Expr:  lit,
	}, nil
}
