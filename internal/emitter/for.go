package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeForInit(forStmt *parser.ForStmt) {}

func (e *Emitter) AfterForInit(forStmt *parser.ForStmt) {
	e.ch.emit(JumpTo)
	jumpTo := e.ch.reserveUint32() // jump to condition

	e.loopStack.push(loop{
		start:     e.ch.ip,
		condStart: jumpTo,
	})
}

func (e *Emitter) BeforeForInc(forStmt *parser.ForStmt) {}

func (e *Emitter) AfterForInc(forStmt *parser.ForStmt) {}

func (e *Emitter) BeforeForCond(forStmt *parser.ForStmt) {
	if top, ok := e.loopStack.top(); ok {
		e.ch.patchUint32(top.condStart, e.ch.ip)
	}
}

func (e *Emitter) AfterForCond(forStmt *parser.ForStmt) {
	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, jumpTo)
	}
}

func (e *Emitter) AfterForBody(forStmt *parser.ForStmt) {
	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
