package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *Emitter) BeforeBlockStmt(block *parser.Block) {
	// e.locals = newLocals(e.locals)
}

func (e *Emitter) AfterBlockStmt(block *parser.Block) {
	// e.locals = e.locals.parent
}
