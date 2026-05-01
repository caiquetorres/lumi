package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeCallExpr(expr *parser.CallExpr) {}

func (e *emitter) AfterCallExpr(call *parser.CallExpr) {
	e.ch.emit(Call)
	e.ch.emitUint8(uint8(len(call.Args)))
}
