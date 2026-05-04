package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeLoopBody(*parser.Loop) {
	e.loopStack.push(loop{
		start: e.ch.ip,
	})
}

func (e *Emitter) AfterLoopBody(*parser.Loop) {
	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
