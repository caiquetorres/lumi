package parser

import (
	"bufio"
	"io"
	"log"

	"github.com/caiquetorres/lumi/internal/lexer"
)

type debugVisitor struct {
	l *lexer.Lexer
	w *bufio.Writer
}

func DebugAst(ast *Ast, l *lexer.Lexer, w io.Writer) error {
	d := &debugVisitor{
		l: l,
		w: bufio.NewWriter(w),
	}
	return Walk(d, ast)
}

func (d *debugVisitor) BeforeAst(ast *Ast) error {
	return nil
}

func (d *debugVisitor) BeforeFunDecl(fd *FunDecl) error {
	mustWrite(d.w.WriteString("fun "))

	id := d.l.Lexeme(fd.Identifier)
	mustWrite(d.w.WriteString(id))

	mustWrite(d.w.WriteRune(' '))
	mustWrite(d.w.WriteString("()"))
	mustWrite(d.w.WriteRune('\n'))

	return d.w.Flush()
}

func (d *debugVisitor) AfterFunDecl(fd *FunDecl) error {
	return nil
}

func (d *debugVisitor) BeforeLiteralExpr(expr *LiteralExpr) error {
	id := d.l.Lexeme(expr.Value)
	mustWrite(d.w.WriteString(id))

	return d.w.Flush()
}

func (d *debugVisitor) BeforeIdentifierExpr(expr *IdentifierExpr) error {
	id := d.l.Lexeme(expr.Name)
	mustWrite(d.w.WriteString(id))

	return d.w.Flush()
}

func (d *debugVisitor) AfterCallExpr(expr *CallExpr) error {
	mustWrite(d.w.WriteString("call "))

	return d.w.Flush()
}

func (d *debugVisitor) AfterStmt(Stmt) error {
	return nil
}

var _ Visitor = (*debugVisitor)(nil)

func mustWrite(_ int, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
