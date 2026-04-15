package emitter

import (
	"bufio"
	"bytes"
	"encoding/binary"
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
	e.ptr++ // 1 byte for the uint8

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
