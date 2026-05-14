package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitFor(forStmt *semantic.For) {
	if forStmt.Init != nil {
		e.emitStmt(forStmt.Init)
	}

	jumpTo := uint32(0)

	if forStmt.Inc != nil {
		e.ch.emit(JumpTo)
		jumpTo = e.ch.reserveUint32() // jump to condition
	}

	e.loopStack.push(loop{
		start:     e.ch.ip,
		condStart: jumpTo,
	})

	if forStmt.Inc != nil {
		e.emitStmt(forStmt.Inc)

		if top, ok := e.loopStack.top(); ok {
			e.ch.patchUint32(top.condStart, e.ch.ip)
		}
	}

	if forStmt.Cond != nil {
		e.emitExpr(forStmt.Cond)

		e.ch.emit(JumpIfFalse)
		jumpTo = e.ch.reserveUint32()

		if top, ok := e.loopStack.top(); ok {
			top.end = append(top.end, jumpTo)
		}
	}

	e.emitBlock(forStmt.Body)

	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
