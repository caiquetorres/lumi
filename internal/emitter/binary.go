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
	}
}
