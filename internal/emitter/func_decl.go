package emitter

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *emitter) AfterFunDecl(*parser.FunDecl) error {
	e.ch.emit(EndScope)
	e.ch.emit(End)

	return nil
}

func (e *emitter) BeforeFunDecl(fn *parser.FunDecl) error {
	e.ch.emit(FnDecl)

	fnName := e.lex.Lexeme(fn.Identifier)
	idx := e.ch.pool.internConstant(fnName)
	e.ch.emitUint32(idx)

	paramCount := len(fn.Params)
	if paramCount > 255 {
		return fmt.Errorf("function '%s' has too many parameters: %d (max 255)", fnName, paramCount)
	}

	e.ch.emitUint8(uint8(paramCount))

	for _, param := range fn.Params {
		paramName := e.lex.Lexeme(param.Name)
		idx := e.ch.pool.internConstant(paramName)
		e.ch.emitUint32(idx)
	}

	offset := e.ch.emitUint32(0)
	entryPoint := e.ch.ip
	e.ch.patchUint32(offset, entryPoint)

	return nil
}

func (e *emitter) AfterParam(*parser.Param) error {
	return nil
}

func (e *emitter) BeforeParam(*parser.Param) error {
	return nil
}
