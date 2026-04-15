package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type Ast struct {
	Statements []TopLevelStmt
}

func (p *Parser) Parse() (*Ast, error) {
	ast := Ast{}

	for !p.peekIs(token.EOF) {
		stmt, err := p.parseTopLevelStmt()
		if err != nil {
			return nil, err
		}

		ast.Statements = append(ast.Statements, stmt)
	}

	return &ast, nil
}
