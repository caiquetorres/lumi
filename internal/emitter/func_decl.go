package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *Emitter) AfterFunDecl(*parser.FunDecl) {
	e.ch.emit(Return)
}

func (e *Emitter) BeforeFunDecl(fn *parser.FunDecl) {
	fnName := e.lex.Lexeme(fn.Identifier)

	fnID := e.fnIDs[fnName]
	e.ch.fnTable[fnID] = e.ch.ip

	if fnName == "main" {
		e.ch.entryPoint = e.ch.ip
		e.ch.hasEntryPoint = true
	}
}

func (e *Emitter) AfterParam(*parser.Param)  {}
func (e *Emitter) BeforeParam(*parser.Param) {}
