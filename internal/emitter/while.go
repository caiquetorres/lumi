package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeWhileCondition(*parser.WhileStmt) {
	e.loopStack.push(loop{
		start: e.ch.ip,
	})
}

func (e *Emitter) AfterWhileCondition(*parser.WhileStmt) {
	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, jumpTo)
	}
}

func (e *Emitter) AfterWhileBody(*parser.WhileStmt) {
	if top, ok := e.loopStack.pop(); ok {
		e.ch.emit(JumpTo)
		e.ch.emitUint32(top.start)

		for _, patch := range top.end {
			e.ch.patchUint32(patch, e.ch.ip)
		}
	}
}
