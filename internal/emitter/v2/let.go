package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) emitLet(l *parser.VarDecl) {
	for _, as := range l.Assignments {
		e.emitExpr(as.Expr)

		name := e.lex.Lexeme(as.Identifier)
		e.storeLocal(name)
	}
}
