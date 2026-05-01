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

		e.emitConst(value)

	case parser.LiteralTrue, parser.LiteralFalse:
		value := lit.Kind == parser.LiteralTrue

		e.emitConst(value)

	case parser.LiteralInt:
		value, err := strconv.Atoi(litValue)
		if err != nil {
			e.setErr(err)
			return
		}

		e.emitConst(value)
	}
}

func (e *emitter) emitConst(value any) {
	e.ch.emit(LoadConst)
	idx := e.ch.pool.InternConstant(value)
	e.ch.emitUint32(idx)
}
