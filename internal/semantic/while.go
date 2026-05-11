package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type WhileStmt struct {
	Condition Expr
	Body      *Block
}

func (t *TypeChecker) analyzeWhileStmt(ws *parser.WhileStmt) *WhileStmt {
	return &WhileStmt{
		Condition: t.analyzeExpr(ws.Condition),
		Body:      t.analyzeBlock(ws.Body),
	}
}
