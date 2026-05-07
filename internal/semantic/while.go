package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type WhileStmt struct {
	Condition Expr
	Body      *Block
}

func whileStmt(ws *parser.WhileStmt) *WhileStmt {
	return &WhileStmt{
		Condition: exprN(ws.Condition),
		Body:      block(ws.Body),
	}
}
