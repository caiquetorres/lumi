package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Ast struct {
	Statements []TopLevelStmt
}

func (p *Parser) Parse() (*Ast, error) {
	ast := Ast{}

	for {
		p.skipWhitespace()

		if p.lookahead().peek().is(token.EOF) {
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
