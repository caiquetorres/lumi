package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitReturn(re *semantic.ReturnStmt) {
	if re.Expr != nil {
		e.emitExpr(re.Expr)
	} else {
		e.ch.emit(PushInt)
		e.ch.emitUint32(0)
	}

	e.ch.emit(Return)
}
