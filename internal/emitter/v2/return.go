package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) emitReturn(re *parser.ReturnStmt) {
	if re.Expr != nil {
		e.emitExpr(re.Expr)
	} else {
		e.ch.emit(PushInt)
		e.ch.emitUint32(0)
	}

	e.ch.emit(Return)
}
