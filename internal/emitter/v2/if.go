package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) emitIf(ifStmt *parser.IfStmt) {
	e.emitExpr(ifStmt.Condition)

	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()
	e.jumpStack.push(jumpTo)

	e.emitBlock(ifStmt.Then)

	jumpTo, _ = e.jumpStack.pop()

	if ifStmt.Else != nil {
		e.ch.emit(JumpTo)
		elseJump := e.ch.reserveUint32()
		e.jumpStack.push(elseJump)
	}

	e.ch.patchUint32(jumpTo, e.ch.ip)

	if ifStmt.Else != nil {
		e.emitBlock(ifStmt.Else)

		elseJump, _ := e.jumpStack.pop()
		e.ch.patchUint32(elseJump, e.ch.ip)
	}
}
