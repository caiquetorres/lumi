package parser

func (p *Parser) parseStmt() (Stmt, error) {
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expectEndOfLine(); err != nil {
		return nil, err
	}

	return expr, nil
}
