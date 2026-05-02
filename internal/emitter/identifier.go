package emitter

import "github.com/caiquetorres/lumi/internal/parser"

// REVIEW: The order of execution is completely wrong

func (e *Emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) {
	name := e.lex.Lexeme(id.Name)

	if fnID, exists := e.funcIDs[name]; exists {
		e.ch.emit(PushFn)
		e.ch.emitUint32(fnID)

		return
	}

	if _, exists := e.nativeFns[name]; exists {
		e.ch.emit(PushNativeFn)

		idx := e.ch.pool.InternConstant(name) // idempotent
		e.ch.emitUint32(idx)

		return
	}

	e.loadLocal(name)
}

func (e *Emitter) loadLocal(name string) {
	sym, exists := e.locals.lookup(name)
	if !exists {
		return
	}

	e.ch.emit(LoadLocal)
	e.ch.emitUint32(uint32(sym.offset))
}
