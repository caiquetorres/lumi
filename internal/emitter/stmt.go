package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) AfterStmt(stmt parser.Stmt) error {
	if _, ok := stmt.(parser.Expr); ok {
		e.ch.emit(Pop)
	}

	return nil
}

func (e *emitter) BeforeReturnStmt(ret *parser.Return) error {
	return nil
}

func (e *emitter) AfterReturnStmt(*parser.Return) error {
	e.ch.emit(Return)

	return nil
}
