package parser

import "github.com/caiquetorres/lumi/internal/token"

type ForStmt struct {
	Identifier token.Token
	Start      Expr
	Op         token.Token // either .. or ..=
	End        Expr
	Block      *Block
}

func (p *Parser) parseFor() (*ForStmt, error) {
	// for <identifier> in <start> .. <end> { ... }
	// for <identifier> in <start> ..= <end> { ... }

	_, err := p.lookahead().next().expect(token.For)
	if err != nil {
		return nil, err
	}

	identifier, err := p.lookahead().next().expect(token.Identifier)
	if err != nil {
		return nil, err
	}

	_, err = p.lookahead().next().expect(token.In)
	if err != nil {
		return nil, err
	}

	start, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	op, err := p.lookahead().next().expectOneOf(token.DotDot, token.DotDotEqual)
	if err != nil {
		return nil, err
	}

	end, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	block, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	return &ForStmt{
		Identifier: identifier,
		Start:      start,
		Op:         op,
		End:        end,
		Block:      block,
	}, nil
}
