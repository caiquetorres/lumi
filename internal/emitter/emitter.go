package emitter

import (
	"io"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

func Emit(ast *parser.Ast, l *lexer.Lexer, w io.Writer) (*Chunk, error) {
	e := newEmitter(l)

	if err := parser.Walk(e, ast); err != nil {
		return nil, err
	}

	builder := newBuilder(w)

	return e.ch, builder.build(e.ch.pool.serialize(), e.ch.code)
}

type blockContext struct {
	breakPatches []uint32
}

type emitter struct {
	ch  *Chunk
	lex *lexer.Lexer

	blockStack []blockContext
	jumpStack  []uint32
}

func newEmitter(lex *lexer.Lexer) *emitter {
	return &emitter{
		ch:        newChunk(),
		lex:       lex,
		jumpStack: make([]uint32, 0),
	}
}

func (e *emitter) BeforeAst(*parser.Ast) error {
	return nil
}

var _ parser.Visitor = (*emitter)(nil)
