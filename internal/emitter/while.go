package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeWhileCondition(*parser.While) error {
	e.loopStack = append(e.loopStack, loop{
		start: e.ch.ip,
	})

	return nil
}

func (e *emitter) AfterWhileCondition(*parser.While) error {
	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.emitUint32(0)

	top := len(e.loopStack) - 1
	e.loopStack[top].end = append(e.loopStack[top].end, jumpTo)

	return nil
}

func (e *emitter) AfterWhileBody(*parser.While) error {
	top := len(e.loopStack) - 1
	jumpTo := e.loopStack[top]
	e.loopStack = e.loopStack[:top]

	e.ch.emit(JumpTo)
	e.ch.emitUint32(jumpTo.start)

	for _, patch := range jumpTo.end {
		e.ch.patchUint32(patch, e.ch.ip)
	}

	return nil
}
