package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeReturnStmt(ret *parser.Return) {}

func (e *emitter) AfterReturnStmt(*parser.Return) {
	e.ch.emit(Return)
}
