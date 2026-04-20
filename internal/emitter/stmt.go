package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) AfterStmt(_ parser.Stmt) error {
	e.ch.emit(Pop)

	return nil
}

func (e *emitter) BeforeReturnStmt(ret *parser.Return) error {
	return nil
}

func (e *emitter) AfterReturnStmt(*parser.Return) error {
	e.ch.emit(Return)
	e.ch.emit(End)

	return nil
}

func (e *emitter) BeforeBreakStmt(*parser.Break) error {
	return nil
}

func (e *emitter) AfterBreakStmt(brk *parser.Break) error {
	e.ch.emit(Return)

	return nil
}

func (e *emitter) AfterVarDecl(vd *parser.VarDecl) error {
	e.ch.emit(VarDecl)

	name := e.lex.Lexeme(vd.Identifier)
	idx := e.ch.pool.internConstant(name)
	e.ch.emitUint32(idx)

	return nil
}

func (e *emitter) BeforeVarDecl(*parser.VarDecl) error {
	return nil
}
