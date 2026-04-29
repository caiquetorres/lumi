package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) AfterIfCondition(ifStmt *parser.If) error {
	e.ch.emit(JumpIfFalse)
	jumpTo := e.ch.emitUint32(0)
	e.jumpStack = append(e.jumpStack, jumpTo)

	return nil
}

func (e *emitter) AfterIfThenBlock(ifStmt *parser.If) error {
	n := len(e.jumpStack) - 1
	jumpTo := e.jumpStack[n]
	e.jumpStack = e.jumpStack[:n]

	if ifStmt.Else != nil {
		e.ch.emit(JumpTo)
		elseJump := e.ch.emitUint32(0)
		e.jumpStack = append(e.jumpStack, elseJump)
	}

	e.ch.patchUint32(jumpTo, e.ch.ip)

	return nil
}

func (e *emitter) AfterElseBlock(ifStmt *parser.If) error {
	n := len(e.jumpStack) - 1
	elseJump := e.jumpStack[n]
	e.jumpStack = e.jumpStack[:n]

	e.ch.patchUint32(elseJump, e.ch.ip)

	return nil
}
