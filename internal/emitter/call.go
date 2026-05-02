package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeCallExpr(expr *parser.CallExpr) {}

func (e *Emitter) AfterCallExpr(call *parser.CallExpr) {
	e.ch.emit(Call)
	e.ch.emitUint8(uint8(len(call.Args)))
}
