package emitter

import (
	"strconv"

	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *emitter) BeforeLiteralExpr(lit *parser.LiteralExpr) {
	litValue := e.lex.Lexeme(lit.Value)

	switch lit.Kind {
	case parser.LiteralString:
		value, err := strconv.Unquote(litValue)
		if err != nil {
			e.setErr(err)
			return
		}

		e.ch.emit(LoadConst)
		idx := e.ch.pool.InternConstant(value)
		e.ch.emitUint32(idx)

	case parser.LiteralTrue:
		e.ch.emit(PushTrue)

	case parser.LiteralFalse:
		e.ch.emit(PushFalse)

	case parser.LiteralInt:
		value, err := strconv.Atoi(litValue)
		if err != nil {
			e.setErr(err)
			return
		}

		e.ch.emit(PushInt)
		e.ch.emitUint32(uint32(value))
	}
}
