package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type While struct {
	Condition Expr
	Body      *Block
}

func (t *TypeChecker) analyzeWhileStmt(ws *parser.WhileStmt) *While {
	return &While{
		Condition: t.analyzeExpr(ws.Condition),
		Body:      t.analyzeBlock(ws.Body),
	}
}
