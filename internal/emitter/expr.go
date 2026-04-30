package emitter

import (
	"strconv"

	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *emitter) BeforeLiteralExpr(lit *parser.LiteralExpr) error {
	litValue := e.lex.Lexeme(lit.Value)

	switch lit.Kind {
	case parser.LiteralString:
		value, err := strconv.Unquote(litValue)
		if err != nil {
			return err
		}

		e.ch.emit(LoadConst)
		idx := e.ch.pool.InternConstant(value)
		e.ch.emitUint32(idx)

	case parser.LiteralTrue, parser.LiteralFalse:
		e.ch.emit(LoadConst)
		value := lit.Kind == parser.LiteralTrue
		idx := e.ch.pool.InternConstant(value)
		e.ch.emitUint32(idx)
	}

	return nil
}

func (e *emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) error {
	e.ch.emit(GetSymbol)

	name := e.lex.Lexeme(id.Name)

	constIdx := e.ch.pool.InternConstant(name)
	e.ch.emitUint32(constIdx)

	return nil
}

func (e *emitter) BeforeBlockExpr(block *parser.Block) error {
	e.ch.emit(BeginScope)

	return nil
}

func (e *emitter) AfterBlockExpr(block *parser.Block) error {
	e.ch.emit(EndScope)

	return nil
}

func (e *emitter) BeforeCallExpr(expr *parser.CallExpr) error {
	return nil
}

func (e *emitter) AfterCallExpr(call *parser.CallExpr) error {
	e.ch.emit(Call)
	e.ch.emitUint8(uint8(len(call.Args)))

	return nil
}
