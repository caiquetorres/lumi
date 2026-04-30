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

	return e.ch, builder.build(e.ch.pool.Serialize(), e.ch.code)
}

type emitter struct {
	ch  *Chunk
	lex *lexer.Lexer

	jumpStack *jumpStack
	loopStack *loopStack
}

func newEmitter(lex *lexer.Lexer) *emitter {
	return &emitter{
		ch:        newChunk(),
		lex:       lex,
		loopStack: newLoopStack(),
		jumpStack: newJumpStack(),
	}
}

func (e *emitter) BeforeAst(*parser.Ast) error {
	return nil
}

var _ parser.Visitor = (*emitter)(nil)

type jumpStack struct {
	data []uint32
}

func newJumpStack() *jumpStack {
	return &jumpStack{
		data: make([]uint32, 0),
	}
}

func (s *jumpStack) push(offset uint32) {
	s.data = append(s.data, offset)
}

func (s *jumpStack) pop() (offset uint32, ok bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	n := len(s.data) - 1
	offset = s.data[n]
	s.data = s.data[:n]
	return offset, true
}

func (s *jumpStack) top() (offset uint32, ok bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	n := len(s.data) - 1
	return s.data[n], true
}

type loop struct {
	start uint32
	end   []uint32
}

type loopStack struct {
	data []loop
}

func newLoopStack() *loopStack {
	return &loopStack{
		data: make([]loop, 0),
	}
}

func (s *loopStack) push(loop loop) {
	s.data = append(s.data, loop)
}

func (s *loopStack) pop() (loop, bool) {
	if len(s.data) == 0 {
		return loop{}, false
	}

	n := len(s.data) - 1
	loop := s.data[n]
	s.data = s.data[:n]
	return loop, true
}

func (s *loopStack) top() (*loop, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	n := len(s.data) - 1
	return &s.data[n], true
}
