package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeContinueStmt(*parser.ContinueStmt) {
	e.ch.emit(JumpTo)

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, e.ch.reserveUint32())
	}
}

func (e *Emitter) AfterContinueStmt(*parser.ContinueStmt) {}
