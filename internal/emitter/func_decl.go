package emitter

import (
	"fmt"

	"github.com/caiquetorres/lumi/internal/parser"
)

func (e *emitter) writeStringConstant(val string) error {
	constIdx := e.pool.internConstant(val)
	return e.writeUint32(constIdx)
}

func (e *emitter) BeforeFunDecl(fn *parser.FunDecl) error {
	if err := e.emit(DeclFun); err != nil {
		return err
	}

	fnName := e.l.Lexeme(fn.Identifier)
	if err := e.writeStringConstant(fnName); err != nil {
		return err
	}

	paramCount := len(fn.Params)
	if paramCount > 255 {
		return fmt.Errorf("function '%s' has too many parameters: %d (max 255)", fnName, paramCount)
	}

	if err := e.writeUint8(uint8(paramCount)); err != nil {
		return err
	}

	for _, param := range fn.Params {
		paramName := e.l.Lexeme(param.Name)
		if err := e.writeStringConstant(paramName); err != nil {
			return err
		}
	}

	// the function body will be emitted after the main code, so we write
	// a placeholder for the function's entry point
	entryPoint := e.ptr + 4
	if err := e.writeUint32(entryPoint); err != nil {
		return err
	}

	if fnName == "main" {
		e.entryPoint = entryPoint
		e.hasEntryPoint = true
	}

	return e.flush()
}

func (e *emitter) AfterFunDecl(fn *parser.FunDecl) error {
	if err := e.emit(End); err != nil {
		return err
	}

	return e.flush()
}
