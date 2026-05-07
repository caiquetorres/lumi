package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) emitStmt(stmt parser.Stmt) {
	switch s := stmt.(type) {
	case *parser.Block:
		e.emitBlock(s)
	case *parser.WhileStmt:
		e.emitWhile(s)
	case *parser.Loop:
		e.emitLoop(s)
	case *parser.ForStmt:
		e.emitFor(s)
	case *parser.IfStmt:
		e.emitIf(s)
	case *parser.BreakStmt:
		e.emitBreak()
	case *parser.VarDecl:
		e.emitLet(s)
	case *parser.ContinueStmt:
		e.emitContinue()
	case *parser.ReturnStmt:
		e.emitReturn(s)
	case parser.Expr:
		e.emitExpr(s)
		e.ch.emit(Pop)
	}
}
