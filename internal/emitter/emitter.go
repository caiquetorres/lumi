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
	var tmp bytes.Buffer

	e := emitter{
		w:          bufio.NewWriter(&tmp),
		l:          l,
		entryPoint: -1,
		pool:       newConstantPool(),
	}

	if err := parser.Walk(&e, ast); err != nil {
		return err
	}

	outFile := bufio.NewWriter(w)

	if _, err := outFile.WriteString("LUMI"); err != nil {
		return err
	}

	// Write the constant pool
	{
		s := e.pool.serialize()
		size := make([]byte, 4)
		binary.BigEndian.PutUint32(size, uint32(len(s)))

		if _, err := outFile.Write(size); err != nil {
			return err
		}

		if _, err := outFile.Write(s); err != nil {
			return err
		}
	}

	if e.entryPoint != -1 {
		outFile.WriteByte(1)

		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(e.entryPoint))

		if _, err := outFile.Write(buf); err != nil {
			return err
		}
	} else {
		outFile.WriteByte(0)
	}

	if _, err := io.Copy(outFile, &tmp); err != nil {
		return err
	}

	return outFile.Flush()
}

type emitter struct {
	ptr        int
	entryPoint int

	w *bufio.Writer

	l    *lexer.Lexer
	pool *constantPool
}

func (e *emitter) VisitAst(*parser.Ast) error {
	return nil
}

func (e *emitter) VisitFunDeclStart(fn *parser.FunDecl) error {
	e.write(DeclFun)

	// load the function's name
	id := e.l.Lexeme(fn.Identifier)
	idx := e.pool.internConstant(id)
	e.writeInt(idx)

	// the function body will be emitted after the main code, so we write
	// a placeholder for the function's entry point
	startPtr := e.ptr + 4
	e.writeInt(startPtr)

	// emit the function body

	if id == "main" {
		e.entryPoint = startPtr
	}

	return e.flush()
}

func (e *emitter) VisitFunDeclEnd(fn *parser.FunDecl) error {
	e.write(End)

	return e.flush()
}

func (e *emitter) VisitLiteralExpr(lit *parser.LiteralExpr) error {
	value := e.l.Lexeme(lit.Value)

	switch lit.Kind {
	case parser.LiteralString:
		value, _ = strconv.Unquote(value)
		e.loadConst(value)
	}

	return e.flush()
}

func (e *emitter) VisitIdentifierExpr(id *parser.IdentifierExpr) error {
	e.write(GetSymbol)

	value := e.l.Lexeme(id.Name)
	idx := e.pool.internConstant(value)
	e.writeInt(idx)

	return e.flush()
}

func (e *emitter) VisitCallExpr(call *parser.CallExpr) error {
	e.write(Call)

	return e.flush()
}

func (e *emitter) VisitStmtEnd(_ parser.Expr) error {
	e.write(Pop)

	return e.flush()
}

var _ parser.Visitor = (*emitter)(nil)

func (e *emitter) flush() error {
	return e.w.Flush()
}

func (e *emitter) write(bytes ...byte) {
	e.ptr += len(bytes)
	_, _ = e.w.Write(bytes)
}

func (e *emitter) loadConst(value any) {
	idx := e.pool.internConstant(value)
	e.writeLoadConst(idx)
}

func (e *emitter) writeLoadConst(idx int) {
	e.write(LoadConst)
	e.writeInt(idx)
}

func (e *emitter) writeInt(value int) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(value))

	e.write(buf...)
}
