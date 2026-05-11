package semantic

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type BinaryExpr struct {
	typedExpr *TypedExpr

	Left     Expr
	Operator token.Token
	Right    Expr
}

func binaryExpr(
	typedInfo *TypedExpr,
	left, right Expr, operator token.Token,
) *BinaryExpr {
	return &BinaryExpr{
		typedExpr: typedInfo,
		Left:      left,
		Operator:  operator,
		Right:     right,
	}
}

func (l *BinaryExpr) Type() *TypedExpr {
	return l.typedExpr
}

var _ Expr = (*BinaryExpr)(nil)

func (t *TypeChecker) analyzeBinaryExpr(be *parser.BinaryExpr) *BinaryExpr {
	var (
		left  = t.analyzeExpr(be.Left)
		right = t.analyzeExpr(be.Right)
	)

	if left.Type().IsAny() || right.Type().IsAny() {
		return binaryExpr(anyExpr(), left, right, be.Operator)
	}

	if left.Type().Kind != right.Type().Kind {
		err := fmt.Errorf("type mismatch: left is %T, right is %T",
			left.Type().Kind, right.Type().Kind)

		t.addErr(err)
		return binaryExpr(anyExpr(), left, right, be.Operator)
	}

	var (
		typedExpr *TypedExpr
		err       error
	)

	if left.Type().IsInt() {
		typedExpr, err = t.analyzeBinaryExprForInts(be, left.Type(), right.Type())
		if err != nil {
			t.addErr(err)
			return binaryExpr(anyExpr(), left, right, be.Operator)
		}
	}

	return binaryExpr(typedExpr, left, right, be.Operator)
}

func (t *TypeChecker) analyzeBinaryExprForInts(
	expr *parser.BinaryExpr,
	left, right *TypedExpr,
) (*TypedExpr, error) {
	if left.IsConst() && right.IsConst() {
		return t.evaluateBinaryForInts(expr.Operator, left, right)
	}

	switch expr.Operator.Kind() {
	case token.Plus, token.Minus, token.Star, token.Slash:
		return newTypedExprKindOnly(left.Kind), nil

	case token.EqualEqual, token.BangEqual:
		return newTypedExprKindOnly(Bool{}), nil

	default:
		panic("unreachable")
	}
}

func (t *TypeChecker) evaluateBinaryForInts(op token.Token, left, right *TypedExpr) (*TypedExpr, error) {
	var (
		leftVal  = left.Value.(int)
		rightVal = right.Value.(int)
	)

	switch op.Kind() {
	case token.Plus:
		return newTypedExpr(Int{}, leftVal+rightVal), nil
	case token.Minus:
		return newTypedExpr(Int{}, leftVal-rightVal), nil
	case token.Star:
		return newTypedExpr(Int{}, leftVal*rightVal), nil
	case token.Slash:
		return t.evaluateDivision(leftVal, rightVal)
	case token.EqualEqual:
		return newTypedExpr(Bool{}, leftVal == rightVal), nil
	case token.BangEqual:
		return newTypedExpr(Bool{}, leftVal != rightVal), nil
	default:
		panic("unreachable")
	}
}

func (t *TypeChecker) evaluateDivision(left, right int) (*TypedExpr, error) {
	if right == 0 {
		return nil, fmt.Errorf("division by zero")
	}

	return newTypedExpr(Int{}, left/right), nil
}
