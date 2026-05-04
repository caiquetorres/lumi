package emitter

import (
	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

type fnVisitor struct {
	globals *globals
	nextID  uint32

	lex *lexer.Lexer
}

func (f *fnVisitor) AfterBinaryExpr(*parser.BinaryExpr)          {}
func (f *fnVisitor) AfterBlockStmt(*parser.Block)                {}
func (f *fnVisitor) AfterBreakStmt(*parser.BreakStmt)            {}
func (f *fnVisitor) AfterCallExpr(*parser.CallExpr)              {}
func (f *fnVisitor) AfterContinueStmt(*parser.ContinueStmt)      {}
func (f *fnVisitor) AfterElseBlock(*parser.IfStmt)               {}
func (f *fnVisitor) AfterFunDecl(*parser.FunDecl)                {}
func (f *fnVisitor) AfterIfCondition(*parser.IfStmt)             {}
func (f *fnVisitor) AfterIfThenBlock(*parser.IfStmt)             {}
func (f *fnVisitor) AfterParam(*parser.Param)                    {}
func (f *fnVisitor) AfterReturnStmt(*parser.ReturnStmt)          {}
func (f *fnVisitor) AfterStmt(parser.Stmt)                       {}
func (f *fnVisitor) AfterVarDecl(*parser.VarDecl)                {}
func (f *fnVisitor) AfterWhileBody(*parser.WhileStmt)            {}
func (f *fnVisitor) AfterWhileCondition(*parser.WhileStmt)       {}
func (f *fnVisitor) BeforeAst(*parser.Ast)                       {}
func (f *fnVisitor) BeforeBinaryExpr(*parser.BinaryExpr)         {}
func (f *fnVisitor) BeforeBlockStmt(*parser.Block)               {}
func (f *fnVisitor) BeforeBreakStmt(*parser.BreakStmt)           {}
func (f *fnVisitor) BeforeCallExpr(*parser.CallExpr)             {}
func (f *fnVisitor) BeforeContinueStmt(*parser.ContinueStmt)     {}
func (f *fnVisitor) BeforeIdentifierExpr(*parser.IdentifierExpr) {}
func (f *fnVisitor) BeforeLiteralExpr(*parser.LiteralExpr)       {}
func (f *fnVisitor) BeforeParam(*parser.Param)                   {}
func (f *fnVisitor) BeforeReturnStmt(*parser.ReturnStmt)         {}
func (f *fnVisitor) BeforeVarDecl(*parser.VarDecl)               {}
func (f *fnVisitor) BeforeWhileCondition(*parser.WhileStmt)      {}
func (f *fnVisitor) AfterAssignment(*parser.Assignment)          {}
func (f *fnVisitor) BeforeAssignment(*parser.Assignment)         {}
func (f *fnVisitor) BeforeForInit(*parser.ForStmt)               {}
func (f *fnVisitor) AfterForInit(*parser.ForStmt)                {}
func (f *fnVisitor) BeforeForInc(*parser.ForStmt)                {}
func (f *fnVisitor) AfterForInc(*parser.ForStmt)                 {}
func (f *fnVisitor) BeforeForCond(*parser.ForStmt)               {}
func (f *fnVisitor) AfterForCond(*parser.ForStmt)                {}
func (f *fnVisitor) AfterForBody(*parser.ForStmt)                {}
func (f *fnVisitor) BeforeLoopBody(*parser.Loop)                 {}
func (f *fnVisitor) AfterLoopBody(*parser.Loop)                  {}

func (f *fnVisitor) BeforeFunDecl(fn *parser.FunDecl) {
	fnName := f.lex.Lexeme(fn.Identifier)
	f.globals.define(fnName)
}

var _ parser.Visitor = (*fnVisitor)(nil)
