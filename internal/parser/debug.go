package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/caiquetorres/lumi/internal/lexer"
)

const indentStep = 2

type debugVisitor struct {
	l *lexer.Lexer
	w *bufio.Writer

	currentIndent int
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

func (d *debugVisitor) BeforeVarDecl(vd *VarDecl) error {
	d.writeIndent()

	name := d.l.Lexeme(vd.Identifier)
	str := fmt.Sprintf("var %s\n", name)
	mustWrite(d.w.WriteString(str))

	return d.flush()
}

func (d *debugVisitor) AfterVarDecl(vd *VarDecl) error {
	return nil
}

func (d *debugVisitor) BeforeFunDecl(fd *FunDecl) error {
	str := fmt.Sprintf("fun %s\n", d.l.Lexeme(fd.Identifier))
	mustWrite(d.w.WriteString(str))

	d.indentIn()

	return d.flush()
}

func (d *debugVisitor) AfterFunDecl(fd *FunDecl) error {
	d.indentOut()

	return nil
}

func (d *debugVisitor) AfterParam(param *Param) error {
	d.writeIndent()

	name := d.l.Lexeme(param.Name)
	str := fmt.Sprintf("param %s\n", name)
	mustWrite(d.w.WriteString(str))

	return d.flush()
}

func (d *debugVisitor) BeforeParam(param *Param) error {
	return d.flush()
}

func (d *debugVisitor) BeforeLiteralExpr(expr *LiteralExpr) error {
	d.writeIndent()

	id := d.l.Lexeme(expr.Value)

	mustWrite(d.w.WriteString(id))
	mustWrite(d.w.WriteRune('\n'))

	return d.flush()
}

func (d *debugVisitor) BeforeIdentifierExpr(expr *IdentifierExpr) error {
	d.writeIndent()

	id := d.l.Lexeme(expr.Name)

	mustWrite(d.w.WriteString(id))
	mustWrite(d.w.WriteRune('\n'))

	return d.flush()
}

func (d *debugVisitor) BeforeCallExpr(expr *CallExpr) error {
	d.writeIndent()

	mustWrite(d.w.WriteString("call\n"))

	d.indentIn()

	return d.flush()
}

func (d *debugVisitor) AfterCallExpr(expr *CallExpr) error {
	d.indentOut()

	return d.flush()
}

func (d *debugVisitor) BeforeReturnStmt(*Return) error {
	d.writeIndent()

	mustWrite(d.w.WriteString("return\n"))

	return d.flush()
}

func (d *debugVisitor) AfterReturnStmt(*Return) error {
	return nil
}

func (d *debugVisitor) AfterStmt(Stmt) error {
	return nil
}

var _ Visitor = (*debugVisitor)(nil)

func (d *debugVisitor) flush() error {
	return d.w.Flush()
}

func (d *debugVisitor) writeIndent() {
	mustWrite(d.w.WriteString(strings.Repeat(" ", d.currentIndent)))
}

func (d *debugVisitor) indentIn() {
	d.currentIndent += indentStep
}

func (d *debugVisitor) indentOut() {
	d.currentIndent -= indentStep
}

func mustWrite(_ int, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
