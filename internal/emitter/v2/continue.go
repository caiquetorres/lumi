package emitter

func (e *Emitter) emitContinue() {
	e.ch.emit(JumpTo)

	if top, ok := e.loopStack.top(); ok {
		e.ch.emitUint32(top.start)
	}
}
