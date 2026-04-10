package parser

import (
	"io"

	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/token"
)

type Ast struct {
	l *lexer.Lexer

	Statements []TopLevelStmt
}

func (p *Parser) DebugAst(ast *Ast, w io.Writer) error {
	p.l.DebugTable()

	for _, stmt := range ast.Statements {
		if err := p.debugTopLevelStmt(stmt, w); err != nil {
			return err
		}

		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseAst() (*Ast, error) {
	ast := Ast{}

	for !p.is(token.EOF) {
		stmt, err := p.parseTopLevelStmt()
		if err != nil {
			return nil, err
		}

		ast.Statements = append(ast.Statements, stmt)
	}

	return &ast, nil
}
