package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) AfterVarDecl(vd *parser.Let) {}

func (e *Emitter) BeforeVarDecl(*parser.Let) {}

func (e *Emitter) BeforeAssignment(*parser.Binding) {}

func (e *Emitter) AfterAssignment(as *parser.Binding) {
	name := e.lex.Lexeme(as.Identifier)
	e.storeLocal(name)
}

func (e *Emitter) storeLocal(name string) {
	e.ch.emit(StoreLocal)

	sym := e.locals.define(name)

	// write offset
	e.ch.emitUint32(uint32(sym.offset))
}
