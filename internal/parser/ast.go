package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Ast struct {
	Statements []TopLevelStmt
}

func (p *Parser) Parse() (*Ast, error) {
	ast := Ast{}

	for !p.peek().is(token.EOF) {
		for p.peek().isOneOf(token.Semicolon, token.NewLine) {
			p.bump()
		}

		if p.peek().is(token.EOF) {
			break
		}

		stmt, err := p.parseTopLevelStmt()
		if err != nil {
			return nil, err
		}

		ast.Statements = append(ast.Statements, stmt)
	}

	return &ast, nil
}
