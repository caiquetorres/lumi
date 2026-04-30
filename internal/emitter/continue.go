package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeContinueStmt(*parser.Continue) {
	e.ch.emit(JumpTo)

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, e.ch.reserveUint32())
	}
}

func (e *emitter) AfterContinueStmt(*parser.Continue) {}
