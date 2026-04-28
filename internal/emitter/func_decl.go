package emitter

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *emitter) AfterFunDecl(*parser.FunDecl) error {
	e.ch.emit(Return)
	e.ch.emit(EndScope)

	return nil
}

func (e *emitter) BeforeFunDecl(fn *parser.FunDecl) error {
	e.ch.emit(FnDecl)

	fnName := e.lex.Lexeme(fn.Identifier)
	idx := e.ch.pool.InternConstant(fnName)
	e.ch.emitUint32(idx)

	paramCount := len(fn.Params)
	if paramCount > 255 {
		return fmt.Errorf("function '%s' has too many parameters: %d (max 255)", fnName, paramCount)
	}

	offset := e.ch.emitUint32(0)
	entryPoint := e.ch.ip
	e.ch.patchUint32(offset, entryPoint)

	e.ch.emit(BeginScope)

	for i := len(fn.Params) - 1; i >= 0; i-- {
		e.ch.emit(DefineSymbol)
		param := fn.Params[i]

		paramName := e.lex.Lexeme(param.Name)
		idx := e.ch.pool.InternConstant(paramName)
		e.ch.emitUint32(idx)
	}

	return nil
}

func (e *emitter) AfterParam(*parser.Param) error {
	return nil
}

func (e *emitter) BeforeParam(*parser.Param) error {
	return nil
}
