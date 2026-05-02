package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) AfterStmt(stmt parser.Stmt) {
	if _, ok := stmt.(parser.Expr); ok {
		e.ch.emit(Pop)
	}
}
