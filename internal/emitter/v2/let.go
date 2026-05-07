package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitLet(l *semantic.VarDecl) {
	for _, as := range l.Assignments {
		e.emitExpr(as.Expr)

		name := e.lex.Lexeme(as.Identifier)
		e.storeLocal(name)
	}
}
