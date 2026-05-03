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

func (d *debugVisitor) AfterBinaryExpr(*BinaryExpr) {}

func (d *debugVisitor) BeforeBinaryExpr(*BinaryExpr) {}

func (d *debugVisitor) AfterBlockStmt(*Block) {}

func (d *debugVisitor) AfterBreakStmt(*BreakStmt) {}

func (d *debugVisitor) BeforeBlockStmt(*Block) {}

func (d *debugVisitor) BeforeContinueStmt(*ContinueStmt) {
	d.writeIndent()

	mustWrite(d.w.WriteString("continue\n"))

	_ = d.flush()
}

func (d *debugVisitor) AfterContinueStmt(*ContinueStmt) {}

func (d *debugVisitor) BeforeBreakStmt(*BreakStmt) {}

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

	_ = d.flush()
}

func (d *debugVisitor) BeforeAssignment(as *Assignment) {
	name := d.l.Lexeme(as.Identifier)
	str := fmt.Sprintf("var %s\n", name)
	mustWrite(d.w.WriteString(str))
}

func (d *debugVisitor) AfterAssignment(assignment *Assignment) {}

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

func (d *debugVisitor) AfterIfCondition(ifStmt *IfStmt) {}

func (d *debugVisitor) AfterIfThenBlock(ifStmt *IfStmt) {
	d.indentOut()
}

func (d *debugVisitor) AfterElseBlock(ifStmt *IfStmt) {
	d.indentOut()
}

func (d *debugVisitor) BeforeWhileCondition(*WhileStmt) {
	d.writeIndent()

	mustWrite(d.w.WriteString("while\n"))

	d.indentIn()

	_ = d.flush()
}

func (d *debugVisitor) AfterWhileCondition(whileStmt *WhileStmt) {}

func (d *debugVisitor) AfterWhileBody(whileStmt *WhileStmt) {
	d.indentOut()
}

func (d *debugVisitor) BeforeForStart(*ForStmt) {}
func (d *debugVisitor) AfterForStart(*ForStmt)  {}
func (d *debugVisitor) BeforeForEnd(*ForStmt)   {}
func (d *debugVisitor) AfterForEnd(*ForStmt)    {}
func (d *debugVisitor) AfterForBody(*ForStmt)   {}

func (d *debugVisitor) AfterForCondition(forStmt *ForStmt) {}

func (d *debugVisitor) BeforeReturnStmt(*ReturnStmt) {
	d.writeIndent()

	mustWrite(d.w.WriteString("return\n"))

	_ = d.flush()
}

func (d *debugVisitor) AfterReturnStmt(*ReturnStmt) {}

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
