package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeReturnStmt(ret *parser.ReturnStmt) {}

func (e *Emitter) AfterReturnStmt(*parser.ReturnStmt) {
	e.ch.emit(Return)
}
