package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) AfterVarDecl(vd *parser.VarDecl) {
	e.ch.emit(DefineSymbol)

	name := e.lex.Lexeme(vd.Identifier)
	idx := e.ch.pool.InternConstant(name)
	e.ch.emitUint32(idx)
}

func (e *emitter) BeforeVarDecl(*parser.VarDecl) {}
