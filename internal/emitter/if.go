package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) AfterIfCondition(ifStmt *parser.IfStmt) {
	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()
	e.jumpStack.push(jumpTo)
}

func (e *Emitter) AfterIfThenBlock(ifStmt *parser.IfStmt) {
	jumpTo, _ := e.jumpStack.pop()

	if ifStmt.Else != nil {
		e.ch.emit(JumpTo)
		elseJump := e.ch.reserveUint32()
		e.jumpStack.push(elseJump)
	}

	e.ch.patchUint32(jumpTo, e.ch.ip)
}

func (e *Emitter) AfterElseBlock(ifStmt *parser.IfStmt) {
	elseJump, _ := e.jumpStack.pop()
	e.ch.patchUint32(elseJump, e.ch.ip)
}
