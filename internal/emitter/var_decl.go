package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) AfterVarDecl(vd *parser.VarDecl) error {
	e.ch.emit(DefineSymbol)

	name := e.lex.Lexeme(vd.Identifier)
	idx := e.ch.pool.internConstant(name)
	e.ch.emitUint32(idx)

	return nil
}

func (e *emitter) BeforeVarDecl(*parser.VarDecl) error {
	return nil
}
