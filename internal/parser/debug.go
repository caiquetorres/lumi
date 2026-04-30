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

func (d *debugVisitor) AfterBlockExpr(*Block) {}

func (d *debugVisitor) AfterBreakStmt(*Break) {}

func (d *debugVisitor) BeforeBlockExpr(*Block) {}

func (d *debugVisitor) BeforeContinueStmt(*Continue) {
	d.writeIndent()

	mustWrite(d.w.WriteString("continue\n"))

	_ = d.flush()
}

func (d *debugVisitor) AfterContinueStmt(*Continue) {}

func (d *debugVisitor) BeforeBreakStmt(*Break) {}

func DebugAst(ast *Ast, l *lexer.Lexer, w io.Writer) {
	d := &debugVisitor{
		l: l,
		w: bufio.NewWriter(w),
	}

	Walk(d, ast)
}

func (d *debugVisitor) BeforeAst(ast *Ast) {}

func (d *debugVisitor) BeforeVarDecl(vd *VarDecl) {
	d.writeIndent()

	name := d.l.Lexeme(vd.Identifier)
	str := fmt.Sprintf("var %s\n", name)
	mustWrite(d.w.WriteString(str))

	_ = d.flush()
}

func (d *debugVisitor) AfterVarDecl(vd *VarDecl) {}

func (d *debugVisitor) BeforeFunDecl(fd *FunDecl) {
	str := fmt.Sprintf("fun %s\n", d.l.Lexeme(fd.Identifier))
	mustWrite(d.w.WriteString(str))

	d.indentIn()

	_ = d.flush()
}

func (d *debugVisitor) AfterFunDecl(fd *FunDecl) {
	d.indentOut()
}

func (d *debugVisitor) AfterParam(param *Param) {
	d.writeIndent()

	name := d.l.Lexeme(param.Name)
	str := fmt.Sprintf("param %s\n", name)
	mustWrite(d.w.WriteString(str))

	_ = d.flush()
}

func (d *debugVisitor) BeforeParam(param *Param) {
	_ = d.flush()
}

func (d *debugVisitor) BeforeLiteralExpr(expr *LiteralExpr) {
	d.writeIndent()

	id := d.l.Lexeme(expr.Value)

	mustWrite(d.w.WriteString(id))
	mustWrite(d.w.WriteRune('\n'))

	_ = d.flush()
}

func (d *debugVisitor) BeforeIdentifierExpr(expr *IdentifierExpr) {
	d.writeIndent()

	id := d.l.Lexeme(expr.Name)

	mustWrite(d.w.WriteString(id))
	mustWrite(d.w.WriteRune('\n'))

	_ = d.flush()
}

func (d *debugVisitor) BeforeCallExpr(expr *CallExpr) {
	d.writeIndent()

	mustWrite(d.w.WriteString("call\n"))

	d.indentIn()

	_ = d.flush()
}

func (d *debugVisitor) AfterCallExpr(expr *CallExpr) {
	d.indentOut()

	_ = d.flush()
}

func (d *debugVisitor) AfterIfCondition(ifStmt *If) {}

func (d *debugVisitor) AfterIfThenBlock(ifStmt *If) {
	d.indentOut()
}

func (d *debugVisitor) AfterElseBlock(ifStmt *If) {
	d.indentOut()
}

func (d *debugVisitor) BeforeWhileCondition(*While) {
	d.writeIndent()

	mustWrite(d.w.WriteString("while\n"))

	d.indentIn()

	_ = d.flush()
}

func (d *debugVisitor) AfterWhileCondition(whileStmt *While) {}

func (d *debugVisitor) AfterWhileBody(whileStmt *While) {
	d.indentOut()
}

func (d *debugVisitor) BeforeReturnStmt(*Return) {
	d.writeIndent()

	mustWrite(d.w.WriteString("return\n"))

	_ = d.flush()
}

func (d *debugVisitor) AfterReturnStmt(*Return) {}

func (d *debugVisitor) AfterStmt(Stmt) {}

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
