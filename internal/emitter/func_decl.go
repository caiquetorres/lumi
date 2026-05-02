package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *emitter) AfterFunDecl(*parser.FunDecl) {
	e.ch.emit(Return)
}

func (e *emitter) BeforeFunDecl(fn *parser.FunDecl) {
	fnName := e.lex.Lexeme(fn.Identifier)
	fnID := e.fnIDs[fnName]
	e.ch.fnTable[fnID] = e.ch.ip

	if fnName == "main" {
		e.ch.entryPoint = e.ch.ip
		e.ch.hasEntryPoint = true
	}
}

func (e *emitter) AfterParam(*parser.Param)  {}
func (e *emitter) BeforeParam(*parser.Param) {}
