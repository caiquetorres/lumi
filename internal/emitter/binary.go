package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

func (e *emitter) BeforeBinaryExpr(be *parser.BinaryExpr) {}

func (e *emitter) AfterBinaryExpr(be *parser.BinaryExpr) {
	switch be.Operator.Kind() {
	case token.Plus:
		e.ch.emit(Add)
	case token.Minus:
		e.ch.emit(Sub)
	case token.Star:
		e.ch.emit(Mul)
	case token.Slash:
		e.ch.emit(Div)

	case token.EqualEqual:
		e.ch.emit(Eq)
	case token.BangEqual:
		e.ch.emit(Eq)
		e.ch.emit(Not)

	case token.Less:
		e.ch.emit(Less)
	case token.LessEqual:
		e.ch.emit(LessEq)
	case token.Greater:
		e.ch.emit(LessEq)
		e.ch.emit(Not)
	case token.GreaterEqual:
		e.ch.emit(Less)
		e.ch.emit(Not)

	case token.Equal:
		e.handleAssignment(be)
	}
}

func (e *emitter) handleAssignment(be *parser.BinaryExpr) {
	switch left := be.Left.(type) {
	case *parser.IdentifierExpr:
		e.ch.emit(SetSymbol)

		name := e.lex.Lexeme(left.Name)
		idx := e.ch.pool.InternConstant(name)
		e.ch.emitUint32(idx)

	default:
		panic("Invalid assignment target.")
	}
}
