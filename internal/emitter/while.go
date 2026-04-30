package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeWhileCondition(*parser.While) error {
	e.loopStack.push(loop{
		start: e.ch.ip,
	})

	return nil
}

func (e *emitter) AfterWhileCondition(*parser.While) error {
	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, jumpTo)
	}

	return nil
}

func (e *emitter) AfterWhileBody(*parser.While) error {
	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}

	return nil
}
