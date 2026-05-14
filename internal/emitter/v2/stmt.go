package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitStmt(stmt semantic.Stmt) {
	switch s := stmt.(type) {
	case *semantic.Block:
		e.emitBlock(s)
	case *semantic.While:
		e.emitWhile(s)
	case *semantic.Loop:
		e.emitLoop(s)
	case *semantic.For:
		e.emitFor(s)
	case *semantic.If:
		e.emitIf(s)
	case *semantic.Break:
		e.emitBreak()
	case *semantic.Let:
		e.emitLet(s)
	case *semantic.Continue:
		e.emitContinue()
	case *semantic.ReturnStmt:
		e.emitReturn(s)
	case semantic.Expr:
		e.emitExpr(s)
		e.ch.emit(Pop)
	}
}
