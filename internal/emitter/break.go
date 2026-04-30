package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeBreakStmt(*parser.Break) error {
	e.ch.emit(JumpTo)
	placeholder := e.ch.emitUint32(0)

	top := len(e.loopStack) - 1
	e.loopStack[top].end = append(e.loopStack[top].end, placeholder)

	return nil
}

func (e *emitter) AfterBreakStmt(brk *parser.Break) error {
	return nil
}
