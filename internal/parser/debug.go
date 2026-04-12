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

func (d *debugVisitor) VisitAst(ast *Ast) error {
	return nil
}

func (d *debugVisitor) VisitFunDeclStart(fd *FunDecl) error {
	mustWrite(d.w.WriteString("fun "))

	id := d.l.Lexeme(fd.Identifier)
	mustWrite(d.w.WriteString(id))

	mustWrite(d.w.WriteRune(' '))
	mustWrite(d.w.WriteString("()"))
	mustWrite(d.w.WriteRune('\n'))

	return d.w.Flush()
}

func (d *debugVisitor) VisitFunDeclEnd(fd *FunDecl) error {
	return nil
}

func (d *debugVisitor) VisitLiteralExpr(expr *LiteralExpr) error {
	id := d.l.Lexeme(expr.Value)
	mustWrite(d.w.WriteString(id))

	return d.w.Flush()
}

func (d *debugVisitor) VisitIdentifierExpr(expr *IdentifierExpr) error {
	id := d.l.Lexeme(expr.Name)
	mustWrite(d.w.WriteString(id))

	return d.w.Flush()
}

func (d *debugVisitor) VisitCallExpr(expr *CallExpr) error {
	mustWrite(d.w.WriteString("call "))

	return d.w.Flush()
}

func (d *debugVisitor) VisitStmtEnd(Expr) error {
	return nil
}

var _ Visitor = (*debugVisitor)(nil)

func mustWrite(_ int, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
