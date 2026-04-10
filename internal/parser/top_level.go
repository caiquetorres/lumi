package parser

import (
	"github.com/caiquetorres/lumi/internal/token"
)

type TopLevelStmt interface{}

func (p *parser) parseTopLevelStmt() (TopLevelStmt, error) {
	switch {
	case p.is(token.Fun):
		return p.parseFunDecl()
	default:
		if _, err := p.peek(); err != nil {
			return nil, err
		}

		return p.expectOneOf(token.Fun)
	}
}
