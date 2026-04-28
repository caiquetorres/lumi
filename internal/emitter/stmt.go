package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) AfterStmt(_ parser.Stmt) error {
	return nil
}

func (e *emitter) BeforeReturnStmt(ret *parser.Return) error {
	return nil
}

func (e *emitter) AfterReturnStmt(*parser.Return) error {
	e.ch.emit(Return)

	return nil
}

func (e *emitter) BeforeBreakStmt(*parser.Break) error {
	return nil
}

func (e *emitter) AfterBreakStmt(brk *parser.Break) error {
	e.ch.emit(Jump)
	patchOffset := e.ch.emitUint32(0)

	top := len(e.blockStack) - 1
	e.blockStack[top].breakPatches = append(e.blockStack[top].breakPatches, patchOffset)

	return nil
}
