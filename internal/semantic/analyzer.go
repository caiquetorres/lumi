package semantic

import (
	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

type Analyzer struct {
	lex *lexer.Lexer
}

type AnalyzedStmt struct {
	IsReachable bool
	Stmt        parser.Stmt
}

type AnalyzedExpr struct {
	Kind  Kind
	Value any
	Expr  parser.Expr
}

func (t *AnalyzedExpr) IsInt() bool {
	_, ok := t.Kind.(Int)
	return ok
}

func (t *AnalyzedExpr) IsString() bool {
	_, ok := t.Kind.(String)
	return ok
}

func (t *AnalyzedExpr) IsBool() bool {
	_, ok := t.Kind.(Bool)
	return ok
}

func (t *AnalyzedExpr) IsFunction() bool {
	_, ok := t.Kind.(Function)
	return ok
}

func (t *AnalyzedExpr) IsConst() bool {
	return t.Value != nil
}

func (t *AnalyzedExpr) AsInt() int {
	if !t.IsInt() {
		panic("type is not an int")
	}

	if val, ok := t.Value.(int); ok {
		return val
	}

	panic("type is not an int")
}

func (t *AnalyzedExpr) AsString() string {
	if !t.IsString() {
		panic("type is not a string")
	}

	if val, ok := t.Value.(string); ok {
		return val
	}

	panic("type is not a string")
}

func (t *AnalyzedExpr) AsBool() bool {
	if !t.IsBool() {
		panic("type is not a bool")
	}

	if val, ok := t.Value.(bool); ok {
		return val
	}

	panic("type is not a bool")
}

type FunctionType struct{}

func (t *AnalyzedExpr) AsFunction() FunctionType {
	if !t.IsFunction() {
		panic("type is not a function")
	}

	if val, ok := t.Value.(FunctionType); ok {
		return val
	}

	panic("type is not a function")
}
