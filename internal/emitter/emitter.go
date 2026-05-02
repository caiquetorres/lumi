package emitter

import (
	"io"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

func Emit(ast *parser.Ast, lex *lexer.Lexer, w io.Writer) (*Chunk, error) {
	fnIDs := buildFnIDs(lex, ast)

	e := newEmitter(lex, fnIDs)
	parser.Walk(e, ast)

	if e.err != nil {
		return nil, e.err
	}

	builder := newBuilder(w)

	return e.ch, builder.build(e.ch)
}

func buildFnIDs(lex *lexer.Lexer, ast *parser.Ast) map[string]uint32 {
	fnVisitor := &fnVisitor{
		fnIDs: make(map[string]uint32),
		lex:   lex,
	}

	parser.Walk(fnVisitor, ast)

	return fnVisitor.fnIDs
}

type Emitter struct {
	ch  *Chunk
	lex *lexer.Lexer

	jumpStack *jumpStack
	loopStack *loopStack

	fnIDs     map[string]uint32
	nativeFns map[string]struct{}
	symTable  *symbolTable

	err error
}

func newEmitter(lex *lexer.Lexer, fnIDs map[string]uint32) *Emitter {
	return &Emitter{
		ch:    newChunk(),
		lex:   lex,
		fnIDs: fnIDs,

		loopStack: newLoopStack(),
		jumpStack: newJumpStack(),
	}
}

func (e *Emitter) registerNativeFn(name string) {
	e.nativeFns[name] = struct{}{}
}

func (e *Emitter) setErr(err error) {
	if e.err == nil {
		e.err = err
	}
}

func (e *Emitter) BeforeAst(*parser.Ast) {}

var _ parser.Visitor = (*Emitter)(nil)
