package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) emitReturn(re *parser.ReturnStmt) {
	if re.Expr != nil {
		e.emitExpr(re.Expr)
	}

	e.ch.emit(Return)
}
