package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) AfterVarDecl(vd *parser.VarDecl) {}

func (e *Emitter) BeforeVarDecl(*parser.VarDecl) {}

func (e *Emitter) BeforeAssignment(*parser.Assignment) {}

func (e *Emitter) AfterAssignment(as *parser.Assignment) {
	name := e.lex.Lexeme(as.Identifier)
	e.storeLocal(name)
}

func (e *Emitter) storeLocal(name string) {
	e.ch.emit(StoreLocal)

	sym := e.locals.define(name)

	// write offset
	e.ch.emitUint32(uint32(sym.offset))
}
