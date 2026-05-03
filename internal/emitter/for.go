package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

func (e *Emitter) BeforeForStart(forStmt *parser.ForStmt) {}

func (e *Emitter) AfterForStart(forStmt *parser.ForStmt) {
	name := e.lex.Lexeme(forStmt.Identifier)
	e.storeLocal(name)

	e.loopStack.push(loop{
		start: e.ch.ip,
	})

	e.loadLocal(name)
}

func (e *Emitter) BeforeForEnd(forStmt *parser.ForStmt) {}

func (e *Emitter) AfterForEnd(forStmt *parser.ForStmt) {
	switch forStmt.Op.Kind() {
	case token.DotDot:
		e.ch.emit(Less)
	case token.DotDotEqual:
		e.ch.emit(LessEq)
	}

	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, jumpTo)
	}
}

func (e *Emitter) AfterForBody(forStmt *parser.ForStmt) {
	name := e.lex.Lexeme(forStmt.Identifier)
	e.loadLocal(name)

	e.ch.emit(PushInt)
	e.ch.emitUint32(1)

	e.ch.emit(Add)

	e.storeLocal(name)

	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
