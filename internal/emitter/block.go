package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeBlockStmt(block *parser.Block) {
	e.ch.emit(BeginScope)
}

func (e *emitter) AfterBlockStmt(block *parser.Block) {
	e.ch.emit(EndScope)
}
