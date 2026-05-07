package emitter

import "github.com/caiquetorres/lumi/internal/semantic"

func (e *Emitter) emitFunDecl(fn *semantic.FunDecl) {
	e.locals = newLocals(nil)

	fnName := e.lex.Lexeme(fn.Identifier)
	funcID, _ := e.globals.lookup(fnName)
	e.ch.fnTable[funcID] = e.ch.ip

	for _, param := range fn.Params {
		name := e.lex.Lexeme(param.Name)
		e.storeLocal(name)
	}

	if fnName == "main" {
		e.ch.entryPoint = e.ch.ip
		e.ch.hasEntryPoint = true
	}

	for _, stmt := range fn.Body {
		e.emitStmt(stmt)
	}

	e.ch.emit(PushInt)
	e.ch.emitUint32(0)
	e.ch.emit(Return)
}
