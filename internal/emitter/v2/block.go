package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) emitBlock(bl *parser.Block) {
	for _, stmt := range bl.Stmts {
		e.emitStmt(stmt)
	}
}
