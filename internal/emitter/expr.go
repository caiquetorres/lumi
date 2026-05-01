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

	case parser.LiteralTrue, parser.LiteralFalse:
		e.ch.emit(LoadConst)
		value := lit.Kind == parser.LiteralTrue
		idx := e.ch.pool.InternConstant(value)
		e.ch.emitUint32(idx)

	case parser.LiteralInt:
		value, err := strconv.Atoi(litValue)
		if err != nil {
			e.setErr(err)
			return
		}

		e.ch.emit(LoadConst)
		idx := e.ch.pool.InternConstant(value)
		e.ch.emitUint32(idx)
	}
}

func (e *emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) {
	e.ch.emit(GetSymbol)

	name := e.lex.Lexeme(id.Name)

	constIdx := e.ch.pool.InternConstant(name)
	e.ch.emitUint32(constIdx)
}

func (e *emitter) BeforeBlockExpr(block *parser.Block) {
	e.ch.emit(BeginScope)
}

func (e *emitter) AfterBlockExpr(block *parser.Block) {
	e.ch.emit(EndScope)
}

func (e *emitter) BeforeCallExpr(expr *parser.CallExpr) {}

func (e *emitter) AfterCallExpr(call *parser.CallExpr) {
	e.ch.emit(Call)
	e.ch.emitUint8(uint8(len(call.Args)))
}
