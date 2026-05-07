package emitter

func (e *Emitter) storeLocal(name string) {
	e.ch.emit(StoreLocal)
	sym := e.locals.define(name)
	e.ch.emitUint32(uint32(sym.offset))
}

func (e *Emitter) loadLocal(name string) {
	local, exists := e.locals.lookup(name)
	if exists {
		e.ch.emit(LoadLocal)
		e.ch.emitUint32(uint32(local.offset))
	}
}
