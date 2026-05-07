package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitStmt(stmt semantic.Stmt) {
	switch s := stmt.(type) {
	case *semantic.Block:
		e.emitBlock(s)
	case *semantic.WhileStmt:
		e.emitWhile(s)
	case *semantic.Loop:
		e.emitLoop(s)
	case *semantic.ForStmt:
		e.emitFor(s)
	case *semantic.IfStmt:
		e.emitIf(s)
	case *semantic.BreakStmt:
		e.emitBreak()
	case *semantic.VarDecl:
		e.emitLet(s)
	case *semantic.ContinueStmt:
		e.emitContinue()
	case *semantic.ReturnStmt:
		e.emitReturn(s)
	case semantic.Expr:
		e.emitExpr(s)
		e.ch.emit(Pop)
	}
}
