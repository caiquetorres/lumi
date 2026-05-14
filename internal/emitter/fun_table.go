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
func (f *fnVisitor) AfterBreakStmt(*parser.Break)                {}
func (f *fnVisitor) AfterCallExpr(*parser.CallExpr)              {}
func (f *fnVisitor) AfterContinueStmt(*parser.Continue)          {}
func (f *fnVisitor) AfterElseBlock(*parser.If)                   {}
func (f *fnVisitor) AfterFunDecl(*parser.FunDecl)                {}
func (f *fnVisitor) AfterIfCondition(*parser.If)                 {}
func (f *fnVisitor) AfterIfThenBlock(*parser.If)                 {}
func (f *fnVisitor) AfterParam(*parser.Param)                    {}
func (f *fnVisitor) AfterReturnStmt(*parser.Return)              {}
func (f *fnVisitor) AfterStmt(parser.Stmt)                       {}
func (f *fnVisitor) AfterVarDecl(*parser.Let)                    {}
func (f *fnVisitor) AfterWhileBody(*parser.WhileStmt)            {}
func (f *fnVisitor) AfterWhileCondition(*parser.WhileStmt)       {}
func (f *fnVisitor) BeforeAst(*parser.Ast)                       {}
func (f *fnVisitor) BeforeBinaryExpr(*parser.BinaryExpr)         {}
func (f *fnVisitor) BeforeBlockStmt(*parser.Block)               {}
func (f *fnVisitor) BeforeBreakStmt(*parser.Break)               {}
func (f *fnVisitor) BeforeCallExpr(*parser.CallExpr)             {}
func (f *fnVisitor) BeforeContinueStmt(*parser.Continue)         {}
func (f *fnVisitor) BeforeIdentifierExpr(*parser.IdentifierExpr) {}
func (f *fnVisitor) BeforeLiteralExpr(*parser.LiteralExpr)       {}
func (f *fnVisitor) BeforeParam(*parser.Param)                   {}
func (f *fnVisitor) BeforeReturnStmt(*parser.Return)             {}
func (f *fnVisitor) BeforeVarDecl(*parser.Let)                   {}
func (f *fnVisitor) BeforeWhileCondition(*parser.WhileStmt)      {}
func (f *fnVisitor) AfterAssignment(*parser.Binding)             {}
func (f *fnVisitor) BeforeAssignment(*parser.Binding)            {}
func (f *fnVisitor) BeforeForInit(*parser.For)                   {}
func (f *fnVisitor) AfterForInit(*parser.For)                    {}
func (f *fnVisitor) BeforeForInc(*parser.For)                    {}
func (f *fnVisitor) AfterForInc(*parser.For)                     {}
func (f *fnVisitor) BeforeForCond(*parser.For)                   {}
func (f *fnVisitor) AfterForCond(*parser.For)                    {}
func (f *fnVisitor) AfterForBody(*parser.For)                    {}
func (f *fnVisitor) BeforeLoopBody(*parser.Loop)                 {}
func (f *fnVisitor) AfterLoopBody(*parser.Loop)                  {}

func (f *fnVisitor) BeforeFunDecl(fn *parser.FunDecl) {
	fnName := f.lex.Lexeme(fn.Identifier)
	f.globals.define(fnName)
}

var _ parser.Visitor = (*fnVisitor)(nil)
