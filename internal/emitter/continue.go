package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeContinueStmt(*parser.ContinueStmt) {
	e.ch.emit(JumpTo)

	if top, ok := e.loopStack.top(); ok {
		e.ch.emitUint32(top.start)
	}
}

func (e *Emitter) AfterContinueStmt(*parser.ContinueStmt) {}
