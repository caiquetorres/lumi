package emitter

func (e *Emitter) emitBreak() {
	e.ch.emit(JumpTo)
	placeholder := e.ch.reserveUint32()

	if top, ok := e.loopStack.top(); ok {
		top.end = append(top.end, placeholder)
	}
}
