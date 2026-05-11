package semantic

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (a *Analyzer) analyzeBinaryExpr(be *parser.BinaryExpr) *BinaryExpr {
	return &BinaryExpr{
		Left:     a.analyzeExpr(be.Left),
		Operator: be.Operator,
		Right:    a.analyzeExpr(be.Right),
	}
}

func (a *Analyzer) AnalyzeExpr(expr parser.Expr) (*AnalyzedExpr, error) {
	switch e := expr.(type) {
	case *parser.LiteralExpr:
		return a.AnalyzeLiteral(e)
	case *parser.BinaryExpr:
		return a.AnalyzeBinaryExpr(e)
	default:
		// TODO: implement other expressions
		panic("unreachable")
	}
}

func (a *Analyzer) AnalyzeBinaryExpr(expr *parser.BinaryExpr) (*AnalyzedExpr, error) {
	left, err := a.AnalyzeExpr(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := a.AnalyzeExpr(expr.Right)
	if err != nil {
		return nil, err
	}

	if left.Kind != right.Kind {
		return nil, fmt.Errorf("type mismatch: left is %T, right is %T", left.Kind, right.Kind)
	}

	if left.IsInt() { // left and right are both ints
		return a.analyzeBinaryExprForInts(expr, left, right)
	}

	return nil, nil
}

func (a *Analyzer) analyzeBinaryExprForInts(
	expr *parser.BinaryExpr,
	left *AnalyzedExpr,
	right *AnalyzedExpr,
) (*AnalyzedExpr, error) {
	if left.IsConst() && right.IsConst() {
		return a.evaluateBinaryForInts(expr, left, right)
	}

	switch expr.Operator.Kind() {
	case token.Plus, token.Minus, token.Star, token.Slash:
		return &AnalyzedExpr{
			Kind:  left.Kind,
			Value: nil,
			Expr:  expr,
		}, nil

	case token.EqualEqual, token.BangEqual:
		return &AnalyzedExpr{
			Kind:  Bool{},
			Value: nil,
			Expr:  expr,
		}, nil

	default:
		panic("unreachable")
	}
}

func (a *Analyzer) evaluateBinaryForInts(
	expr *parser.BinaryExpr,
	left, right *AnalyzedExpr,
) (*AnalyzedExpr, error) {
	var (
		leftVal  = left.Value.(int)
		rightVal = right.Value.(int)
	)

	switch expr.Operator.Kind() {
	case token.Plus:
		return a.evaluateAddition(expr, leftVal, rightVal)
	case token.Minus:
		return a.evaluateSubtraction(expr, leftVal, rightVal)
	case token.Star:
		return a.evaluateMultiplication(expr, leftVal, rightVal)
	case token.Slash:
		return a.evaluateDivision(expr, leftVal, rightVal)
	case token.EqualEqual:
		return a.evaluateEqual(expr, leftVal, rightVal)
	case token.BangEqual:
		return a.evaluateNotEqual(expr, leftVal, rightVal)
	default:
		panic("unreachable")
	}
}

func (a *Analyzer) evaluateAddition(
	expr *parser.BinaryExpr,
	left, right int,
) (*AnalyzedExpr, error) {
	return &AnalyzedExpr{
		Kind:  Int{},
		Value: left + right,
		Expr:  expr,
	}, nil
}

func (a *Analyzer) evaluateSubtraction(
	expr *parser.BinaryExpr,
	left, right int,
) (*AnalyzedExpr, error) {
	return &AnalyzedExpr{
		Kind:  Int{},
		Value: left - right,
		Expr:  expr,
	}, nil
}

func (a *Analyzer) evaluateMultiplication(
	expr *parser.BinaryExpr,
	left, right int,
) (*AnalyzedExpr, error) {
	return &AnalyzedExpr{
		Kind:  Int{},
		Value: left * right,
		Expr:  expr,
	}, nil
}

func (a *Analyzer) evaluateDivision(
	expr *parser.BinaryExpr,
	left, right int,
) (*AnalyzedExpr, error) {
	if right == 0 {
		return nil, fmt.Errorf("division by zero")
	}

	return &AnalyzedExpr{
		Kind:  Int{},
		Value: left / right,
		Expr:  expr,
	}, nil
}

func (a *Analyzer) evaluateEqual(
	expr *parser.BinaryExpr,
	left, right int,
) (*AnalyzedExpr, error) {
	return &AnalyzedExpr{
		Kind:  Bool{},
		Value: left == right,
		Expr:  expr,
	}, nil
}

func (a *Analyzer) evaluateNotEqual(
	expr *parser.BinaryExpr,
	left, right int,
) (*AnalyzedExpr, error) {
	return &AnalyzedExpr{
		Kind:  Bool{},
		Value: left != right,
		Expr:  expr,
	}, nil
}
