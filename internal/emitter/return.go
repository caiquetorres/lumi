package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeReturnStmt(ret *parser.ReturnStmt) {}

func (e *emitter) AfterReturnStmt(*parser.ReturnStmt) {
	e.ch.emit(Return)
}
