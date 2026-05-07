package emitter

import (
	"strconv"

	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

func (e *Emitter) emitExpr(expr parser.Expr) {
	switch expr := expr.(type) {
	case *parser.LiteralExpr:
		e.emitLiteral(expr)
	case *parser.IdentifierExpr:
		e.emitIdentifier(expr)
	case *parser.CallExpr:
		e.emitCall(expr)
	case *parser.BinaryExpr:
		e.emitBinaryExpr(expr)
	default:
		panic("unreachable")
	}
}

func (e *Emitter) emitLiteral(lit *parser.LiteralExpr) {
	litValue := e.lex.Lexeme(lit.Value)

	switch lit.Kind {
	case parser.LiteralString:
		value, err := strconv.Unquote(litValue)
		if err != nil {
			e.setErr(err)
			return
		}

		e.ch.emit(PushString)
		idx := e.ch.pool.InternConstant(value)
		e.ch.emitUint32(idx)

	case parser.LiteralTrue:
		e.ch.emit(PushTrue)

	case parser.LiteralFalse:
		e.ch.emit(PushFalse)

	case parser.LiteralInt:
		value, err := strconv.Atoi(litValue)
		if err != nil {
			e.setErr(err)
			return
		}

		e.ch.emit(PushInt)
		e.ch.emitUint32(uint32(value))
	}
}

func (e *Emitter) emitIdentifier(id *parser.IdentifierExpr) {
	name := e.lex.Lexeme(id.Name)

	e.loadLocal(name)

	if fnID, exists := e.globals.lookup(name); exists {
		e.ch.emit(PushFn)
		e.ch.emitUint32(fnID)

		return
	}

	if _, exists := e.nativeFns[name]; exists {
		e.ch.emit(PushNativeFn)

		idx := e.ch.pool.InternConstant(name) // idempotent
		e.ch.emitUint32(idx)

		return
	}
}

func (e *Emitter) emitCall(ca *parser.CallExpr) {
	for i := len(ca.Args) - 1; i >= 0; i-- {
		e.emitExpr(ca.Args[i])
	}

	e.emitExpr(ca.Callee)

	e.ch.emit(Call)
	e.ch.emitUint8(uint8(len(ca.Args)))
}

func (e *Emitter) emitBinaryExpr(be *parser.BinaryExpr) {
	e.emitExpr(be.Left)
	e.emitExpr(be.Right)

	switch be.Operator.Kind() {
	case token.Plus:
		e.ch.emit(Add)
	case token.Minus:
		e.ch.emit(Sub)
	case token.Star:
		e.ch.emit(Mul)
	case token.Slash:
		e.ch.emit(Div)

	case token.Equal:
		e.handleAssignment(be)

	case token.PlusEqual:
		e.handleCompoundAssignment(be, Add)
	case token.MinusEqual:
		e.handleCompoundAssignment(be, Sub)
	case token.StarEqual:
		e.handleCompoundAssignment(be, Mul)
	case token.SlashEqual:
		e.handleCompoundAssignment(be, Div)

	case token.EqualEqual:
		e.ch.emit(Eq)
	case token.BangEqual:
		e.ch.emit(Eq)
		e.ch.emit(Not)

	case token.Less:
		e.ch.emit(Less)
	case token.LessEqual:
		e.ch.emit(LessEq)
	case token.Greater:
		e.ch.emit(LessEq)
		e.ch.emit(Not)
	case token.GreaterEqual:
		e.ch.emit(Less)
		e.ch.emit(Not)
	}
}

func (e *Emitter) handleAssignment(be *parser.BinaryExpr) {
	switch left := be.Left.(type) {
	case *parser.IdentifierExpr:
		name := e.lex.Lexeme(left.Name)
		e.storeLocal(name)
		e.loadLocal(name)

	default:
		panic("Invalid assignment target.")
	}
}

func (e *Emitter) handleCompoundAssignment(be *parser.BinaryExpr, op byte) {
	switch left := be.Left.(type) {
	case *parser.IdentifierExpr:
		name := e.lex.Lexeme(left.Name)

		e.ch.emit(op)
		e.storeLocal(name)
		e.loadLocal(name)

	default:
		panic("Invalid assignment target.")
	}
}
