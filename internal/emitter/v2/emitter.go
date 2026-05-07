package emitter

import (
	"github.com/caiquetorres/lumi/internal/lexer"
)

type Emitter struct {
	ch  *Chunk
	lex *lexer.Lexer

	jumpStack *jumpStack
	loopStack *loopStack

	locals  *locals
	globals *globals

	nativeFns map[string]struct{}

	err error
}

func newEmitter(lex *lexer.Lexer, globals *globals) *Emitter {
	return &Emitter{
		ch:        newChunk(len(globals.symbols)),
		lex:       lex,
		globals:   globals,
		nativeFns: make(map[string]struct{}),
		loopStack: newLoopStack(),
		jumpStack: newJumpStack(),
	}
}

func (e *Emitter) setErr(err error) {
	if e.err == nil {
		e.err = err
	}
}
