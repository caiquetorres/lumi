package emitter

import (
	"io"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
)

const lumiMagic = "LUMI"

func WriteLumiFile(ch *Chunk, w io.Writer) error {
	_, err := w.Write([]byte(lumiMagic))
	if err != nil {
		return err
	}

	return ch.Serialize(w)
}

func Emit(ast *parser.Ast, lex *lexer.Lexer) (*Chunk, error) {
	globals := buildGlobals(ast, lex)

	e := newEmitter(lex, globals)
	e.registerNativeFn("println")
	e.registerNativeFn("printf")
	e.registerNativeFn("sprintf")

	e.emitAst(ast)

	if e.err != nil {
		return nil, e.err
	}

	return e.ch, nil
}

func (e *Emitter) emitAst(ast *parser.Ast) {
	for _, stmt := range ast.Statements {
		switch stmt := stmt.(type) {
		case *parser.FunDecl:
			e.emitFunDecl(stmt)
		}
	}
}

func buildGlobals(ast *parser.Ast, lex *lexer.Lexer) *globals {
	g := newGlobals()
	for _, stmt := range ast.Statements {
		switch stmt := stmt.(type) {
		case *parser.FunDecl:
			fnName := lex.Lexeme(stmt.Identifier)
			_ = g.define(fnName)
		}
	}

	return g
}

func (e *Emitter) registerNativeFn(name string) {
	e.nativeFns[name] = struct{}{}
}
