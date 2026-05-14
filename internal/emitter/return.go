package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeReturnStmt(*parser.Return) {}

func (e *Emitter) AfterReturnStmt(ret *parser.Return) {
	if ret.Expr == nil {
		e.ch.emit(PushInt)
		e.ch.emitUint32(0)
	}

	e.ch.emit(Return)
}
