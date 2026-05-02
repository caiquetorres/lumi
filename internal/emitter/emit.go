package emitter

import (
	"io"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

const (
	lumiMagic = "LUMI"
)

func Emit(ast *parser.Ast, lex *lexer.Lexer, w io.Writer) (*Chunk, error) {
	funcIDs := collectFunctionIDs(lex, ast)

	e := newEmitter(lex, funcIDs)
	e.registerNativeFn("println")
	e.registerNativeFn("printf")
	e.registerNativeFn("sprintf")

	parser.Walk(e, ast)

	if e.err != nil {
		return nil, e.err
	}

	if _, err := w.Write([]byte(lumiMagic)); err != nil {
		return nil, err
	}

	return e.ch, e.ch.Serialize(w)
}

func collectFunctionIDs(lex *lexer.Lexer, ast *parser.Ast) map[string]uint32 {
	fnVisitor := &fnVisitor{
		fnIDs: make(map[string]uint32),
		lex:   lex,
	}

	parser.Walk(fnVisitor, ast)

	return fnVisitor.fnIDs
}

func (e *Emitter) registerNativeFn(name string) {
	e.nativeFns[name] = struct{}{}
}
