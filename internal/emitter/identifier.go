package emitter

import "github.com/caiquetorres/lumi/internal/parser"

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
}
