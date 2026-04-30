package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeContinueStmt(*parser.Continue) error {
	e.ch.emit(JumpTo)

	top := len(e.loopStack) - 1
	e.ch.emitUint32(e.loopStack[top].start)

	return nil
}

func (e *emitter) AfterContinueStmt(*parser.Continue) error {
	return nil
}
