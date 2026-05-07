package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitBlock(bl *semantic.Block) {
	for _, stmt := range bl.Stmts {
		e.emitStmt(stmt)
	}
}
