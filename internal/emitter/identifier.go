package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) {
	e.ch.emit(GetSymbol)

	name := e.lex.Lexeme(id.Name)

	constIdx := e.ch.pool.InternConstant(name)
	e.ch.emitUint32(constIdx)
}
