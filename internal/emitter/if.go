package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) AfterIfCondition(ifStmt *parser.If) error {
	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.reserveUint32()
	e.jumpStack.push(jumpTo)

	return nil
}

func (e *emitter) AfterIfThenBlock(ifStmt *parser.If) error {
	jumpTo, _ := e.jumpStack.pop()

	if ifStmt.Else != nil {
		e.ch.emit(JumpTo)
		elseJump := e.ch.reserveUint32()
		e.jumpStack.push(elseJump)
	}

	e.ch.patchUint32(jumpTo, e.ch.ip)

	return nil
}

func (e *emitter) AfterElseBlock(ifStmt *parser.If) error {
	elseJump, _ := e.jumpStack.pop()
	e.ch.patchUint32(elseJump, e.ch.ip)

	return nil
}
