package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) {
	name := e.lex.Lexeme(id.Name)

	if fnID, exists := e.fnIDs[name]; exists {
		e.ch.emit(PushFn)
		e.ch.emitUint32(fnID)
		return
	}
}
