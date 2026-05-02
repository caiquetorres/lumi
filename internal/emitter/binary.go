package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

func (e *Emitter) BeforeBinaryExpr(be *parser.BinaryExpr) {}

func (e *Emitter) AfterBinaryExpr(be *parser.BinaryExpr) {
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

func (e *Emitter) handleAssignment(be *parser.BinaryExpr) {
	switch left := be.Left.(type) {
	case *parser.IdentifierExpr:
		name := e.lex.Lexeme(left.Name)
		e.storeLocal(name)

	default:
		panic("Invalid assignment target.")
	}
}
