package semantic

import (
	"strconv"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

type Analyzer struct {
	lex *lexer.Lexer
}

type semanticError struct {
	errs    []error
	isFatal bool
}

func (e *semanticError) Error() string {
	panic("unimplemented")
}

func Analyze(ast *parser.Ast) error {
	return nil
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

type ExprType any

type (
	ExprInt    struct{}
	ExprString struct{}
	ExprBool   struct{}
)

type AnalyzedStmt struct {
	IsReachable bool
	Stmt        parser.Stmt
}

type AnalyzedExpr struct {
	Type  ExprType
	Value any
	Expr  parser.Expr
}

func (e *AnalyzedExpr) IsConst() bool {
	return e.Value != nil
}

func (e *AnalyzedExpr) ConstValue() (any, bool) {
	if !e.IsConst() {
		return nil, false
	}

	return e.Value, true
}

func (e *AnalyzedExpr) AsInt() (int, bool) {
	if !e.IsInt() {
		return 0, false
	}

	val, ok := e.ConstValue()
	if !ok {
		return 0, false
	}

	intVal, ok := val.(int)
	if !ok {
		return 0, false
	}

	return intVal, true
}

func (e *AnalyzedExpr) IsConstOfType(ty ExprType) bool {
	return e.IsConst() && e.IsType(ty)
}

func (e *AnalyzedExpr) IsBoolConst() bool {
	return e.IsConstOfType(ExprBool{})
}

func (e *AnalyzedExpr) IsInt() bool {
	return e.IsType(ExprInt{})
}

func (e *AnalyzedExpr) IsString() bool {
	return e.IsType(ExprString{})
}

func (e *AnalyzedExpr) IsBool() bool {
	return e.IsType(ExprBool{})
}

func (e *AnalyzedExpr) IsType(ty ExprType) bool {
	return e.Type == ty
}

func (a *Analyzer) analyzeLiteral(lit *parser.LiteralExpr) (*AnalyzedExpr, error) {
	var (
		ty  ExprType
		val any

		err error
	)

	switch lit.Kind {
	case parser.LiteralInt:
		ty = ExprInt{}

		text := a.lex.Lexeme(lit.Value)
		val, err = strconv.Atoi(text)
		if err != nil {
			return nil, err
		}

	case parser.LiteralString:
		ty = ExprString{}

	case parser.LiteralFalse, parser.LiteralTrue:
		ty = ExprBool{}
		text := a.lex.Lexeme(lit.Value)

		val = text == "true"
	default:
		panic("unreachable")
	}

	return &AnalyzedExpr{
		Type:  ty,
		Value: val,
		Expr:  lit,
	}, nil
}
