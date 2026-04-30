package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeReturnStmt(ret *parser.Return) error {
	return nil
}

func (e *emitter) AfterReturnStmt(*parser.Return) error {
	e.ch.emit(Return)

	return nil
}
