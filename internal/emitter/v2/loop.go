package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitLoop(l *semantic.Loop) {
	e.loopStack.push(loop{
		start: e.ch.ip,
	})

	e.emitBlock(l.Body)

	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
