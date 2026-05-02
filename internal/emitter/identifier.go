package emitter

import (
	"github.com/caiquetorres/lumi/internal/parser"
)

// REVIEW: The order of execution is completely wrong

func (e *Emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) {
	name := e.lex.Lexeme(id.Name)

	local, exists := e.locals.lookup(name)
	if exists {
		e.ch.emit(LoadLocal)
		e.ch.emitUint32(uint32(local.offset))

		return
	}

	if fnID, exists := e.globals.lookup(name); exists {
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
}
