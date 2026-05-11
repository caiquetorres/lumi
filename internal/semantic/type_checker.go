package semantic

import "github.com/caiquetorres/lumi/internal/lexer"

type TypeChecker struct {
	lex      *lexer.Lexer
	symTable *SymbolTable

	err []error
}

func NewAnalyzer(lex *lexer.Lexer) *TypeChecker {
	return &TypeChecker{
		lex:      lex,
		symTable: NewRootSymbolTable(),
	}
}

func (t *TypeChecker) addErr(err error) {
	t.err = append(t.err, err)
}

type TypedExpr struct {
	Kind  Kind
	Value any
}

func newTypedExpr(kind Kind, value any) *TypedExpr {
	return &TypedExpr{
		Kind:  kind,
		Value: value,
	}
}

func newTypedExprKindOnly(kind Kind) *TypedExpr {
	return &TypedExpr{
		Kind:  kind,
		Value: nil,
	}
}

func anyExpr() *TypedExpr {
	return newTypedExprKindOnly(Any{})
}

func (t *TypedExpr) IsAny() bool {
	_, ok := t.Kind.(Any)
	return ok
}

func (t *TypedExpr) IsInt() bool {
	_, ok := t.Kind.(Int)
	return ok
}

func (t *TypedExpr) IsString() bool {
	_, ok := t.Kind.(String)
	return ok
}

func (t *TypedExpr) IsBool() bool {
	_, ok := t.Kind.(Bool)
	return ok
}

func (t *TypedExpr) IsFunction() bool {
	_, ok := t.Kind.(Function)
	return ok
}

func (t *TypedExpr) IsConst() bool {
	return t.Value != nil
}

func (t *TypedExpr) AsInt() int {
	if !t.IsInt() {
		panic("type is not an int")
	}

	if val, ok := t.Value.(int); ok {
		return val
	}

	panic("type is not an int")
}

func (t *TypedExpr) AsString() string {
	if !t.IsString() {
		panic("type is not a string")
	}

	if val, ok := t.Value.(string); ok {
		return val
	}

	panic("type is not a string")
}

func (t *TypedExpr) AsBool() bool {
	if !t.IsBool() {
		panic("type is not a bool")
	}

	if val, ok := t.Value.(bool); ok {
		return val
	}

	panic("type is not a bool")
}

func (t *TypedExpr) AsFunction() Function {
	if !t.IsFunction() {
		panic("type is not a function")
	}

	if val, ok := t.Value.(Function); ok {
		return val
	}

	panic("type is not a function")
}
