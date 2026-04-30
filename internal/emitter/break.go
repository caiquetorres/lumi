package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeBreakStmt(*parser.Break) {
	e.ch.emit(JumpTo)
	placeholder := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, placeholder)
	}
}

func (e *emitter) AfterBreakStmt(*parser.Break) {}
