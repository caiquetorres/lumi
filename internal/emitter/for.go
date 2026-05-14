package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeForInit(forStmt *parser.For) {}

func (e *Emitter) AfterForInit(forStmt *parser.For) {}

func (e *Emitter) BeforeForInc(forStmt *parser.For) {
	var jumpTo uint32

	if forStmt.Inc != nil {
		e.ch.emit(JumpTo)
		jumpTo = e.ch.reserveUint32() // jump to condition
	}

	e.loopStack.push(loop{
		start:     e.ch.ip,
		condStart: jumpTo,
	})
}

func (e *Emitter) AfterForInc(forStmt *parser.For) {
	if forStmt.Inc != nil {
		if top, ok := e.loopStack.top(); ok {
			e.ch.patchUint32(top.condStart, e.ch.ip)
		}
	}
}

func (e *Emitter) BeforeForCond(forStmt *parser.For) {}

func (e *Emitter) AfterForCond(forStmt *parser.For) {
	if forStmt.Cond == nil {
		return
	}

	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, jumpTo)
	}
}

func (e *Emitter) AfterForBody(forStmt *parser.For) {
	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
