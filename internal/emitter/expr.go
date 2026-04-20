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
		idx := e.ch.pool.internConstant(value)
		e.ch.emitUint32(idx)
	}

	return nil
}

func (e *emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) error {
	e.ch.emit(GetSymbol)

	name := e.lex.Lexeme(id.Name)

	constIdx := e.ch.pool.internConstant(name)
	e.ch.emitUint32(constIdx)

	return nil
}

func (e *emitter) BeforeBlockExpr(block *parser.BlockExpr) error {
	e.blockStack = append(e.blockStack, blockContext{})
	e.ch.emit(BeginScope)

	return nil
}

func (e *emitter) AfterBlockExpr(block *parser.BlockExpr) error {
	e.ch.emit(EndScope)

	endOffset := e.ch.ip
	top := len(e.blockStack) - 1

	for _, off := range e.blockStack[top].breakPatches {
		e.ch.patchUint32(off, endOffset)
	}

	e.blockStack = e.blockStack[:top]

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
