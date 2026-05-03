package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) AfterVarDecl(vd *parser.VarDecl) {
	name := e.lex.Lexeme(vd.Identifier)
	e.storeLocal(name)
	e.ch.emit(Pop)
}

func (e *Emitter) BeforeVarDecl(*parser.VarDecl) {}

func (e *Emitter) storeLocal(name string) {
	e.ch.emit(StoreLocal)

	sym := e.locals.define(name)

	// write offset
	e.ch.emitUint32(uint32(sym.offset))
}
