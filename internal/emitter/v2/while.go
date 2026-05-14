package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitWhile(wh *semantic.While) {
	e.loopStack.push(loop{
		start: e.ch.ip,
	})

	e.emitExpr(wh.Condition)

	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, jumpTo)
	}

	e.emitBlock(wh.Body)

	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
