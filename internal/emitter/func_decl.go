package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *Emitter) AfterFunDecl(*parser.FunDecl) {
	e.ch.emit(Return)
}

func (e *Emitter) BeforeFunDecl(fn *parser.FunDecl) {
	e.locals = newLocals(nil)

	fnName := e.lex.Lexeme(fn.Identifier)

	funcID, _ := e.globals.lookup(fnName)
	e.ch.fnTable[funcID] = e.ch.ip

	for _, param := range fn.Params {
		name := e.lex.Lexeme(param.Name)
		e.storeLocal(name)
		e.ch.emit(Pop)
	}

	if fnName == "main" {
		e.ch.entryPoint = e.ch.ip
		e.ch.hasEntryPoint = true
	}
}

func (e *Emitter) AfterParam(*parser.Param)  {}
func (e *Emitter) BeforeParam(*parser.Param) {}
