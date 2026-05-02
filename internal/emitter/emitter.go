package emitter

import (
	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

type Emitter struct {
	ch  *Chunk
	lex *lexer.Lexer

	jumpStack *jumpStack
	loopStack *loopStack
	locals    *locals

	funcIDs   map[string]uint32
	nativeFns map[string]struct{}

	err error
}

func newEmitter(lex *lexer.Lexer, fnIDs map[string]uint32) *Emitter {
	return &Emitter{
		ch:  newChunk(len(fnIDs)),
		lex: lex,

		funcIDs:   fnIDs,
		nativeFns: make(map[string]struct{}),

		locals:    newLocals(nil, 0),
		loopStack: newLoopStack(),
		jumpStack: newJumpStack(),
	}
}

func (e *Emitter) setErr(err error) {
	if e.err == nil {
		e.err = err
	}
}

func (e *Emitter) BeforeAst(*parser.Ast) {}

var _ parser.Visitor = (*Emitter)(nil)
