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
}

func newEmitter(lex *lexer.Lexer) *emitter {
	return &emitter{
		ch:  newChunk(),
		lex: lex,
	}
}

func (e *emitter) BeforeAst(*parser.Ast) error {
	return nil
}

var _ parser.Visitor = (*emitter)(nil)

// func formatBytecode(code []byte, w io.Writer) {
// 	if len(code) == 0 {
// 		return
// 	}

// 	for idx, b := range code {
// 		if idx > 0 {
// 			w.Write([]byte{' '})
// 		}

// 		_, _ = fmt.Fprintf(w, "0x%02X", b)
// 	}

// 	fmt.Fprintf(w, "\n")
// }
