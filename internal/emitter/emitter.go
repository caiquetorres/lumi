package emitter

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

func Emit(ast *parser.Ast, l *lexer.Lexer, w io.Writer) error {
	tmp := &bytes.Buffer{}
	e := newEmitter(l, tmp)

	if err := parser.Walk(e, ast); err != nil {
		return err
	}

	builder := newBuilder(w)
	return builder.build(e.pool.serialize(), e.hasEntryPoint, e.entryPoint, tmp)
}

type emitter struct {
	ptr uint32

	entryPoint    uint32
	hasEntryPoint bool

	w    *bufio.Writer
	l    *lexer.Lexer
	pool *constantPool
}

func newEmitter(l *lexer.Lexer, w io.Writer) *emitter {
	return &emitter{
		w:    bufio.NewWriter(w),
		l:    l,
		pool: newConstantPool(),
	}
}

func (e *emitter) BeforeAst(*parser.Ast) error {
	return nil
}

func (e *emitter) BeforeFunDecl(fn *parser.FunDecl) error {
	e.emit(DeclFun)

	fnName := e.l.Lexeme(fn.Identifier)
	constIdx := e.pool.internConstant(fnName)
	if err := e.writeUint32(constIdx); err != nil {
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
		paramIdx := e.pool.internConstant(paramName)
		if err := e.writeUint32(paramIdx); err != nil {
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

func (e *emitter) BeforeLiteralExpr(lit *parser.LiteralExpr) error {
	litValue := e.l.Lexeme(lit.Value)

	switch lit.Kind {
	case parser.LiteralString:
		value, err := strconv.Unquote(litValue)
		if err != nil {
			return err
		}

		constIdx := e.pool.internConstant(value)
		if err := e.emit(LoadConst); err != nil {
			return err
		}
		if err := e.writeUint32(constIdx); err != nil {
			return err
		}
	}

	return e.flush()
}

func (e *emitter) BeforeIdentifierExpr(id *parser.IdentifierExpr) error {
	if err := e.emit(GetSymbol); err != nil {
		return err
	}

	idName := e.l.Lexeme(id.Name)

	constIdx := e.pool.internConstant(idName)
	if err := e.writeUint32(constIdx); err != nil {
		return err
	}

	return e.flush()
}

func (e *emitter) BeforeCallExpr(expr *parser.CallExpr) error {
	return nil
}

func (e *emitter) AfterCallExpr(call *parser.CallExpr) error {
	if err := e.emit(Call); err != nil {
		return err
	}

	if err := e.writeUint8(uint8(len(call.Args))); err != nil {
		return err
	}

	return e.flush()
}

func (e *emitter) AfterStmt(_ parser.Stmt) error {
	if err := e.emit(Pop); err != nil {
		return err
	}

	return e.flush()
}

func (e *emitter) AfterParam(*parser.Param) error {
	return nil
}

func (e *emitter) BeforeParam(*parser.Param) error {
	return nil
}

var _ parser.Visitor = (*emitter)(nil)

func (e *emitter) emit(b byte) error {
	return e.writeUint8(b)
}

func (e *emitter) writeUint8(value uint8) error {
	e.ptr++
	return e.w.WriteByte(value)
}

func (e *emitter) writeUint32(value uint32) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], value)

	e.ptr += 4 // 4 bytes for the uint32

	_, err := e.w.Write(buf[:])
	return err
}

func (e *emitter) flush() error {
	return e.w.Flush()
}
