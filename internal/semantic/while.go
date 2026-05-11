package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type WhileStmt struct {
	Condition Expr
	Body      *Block
}

func (a *TypeChecker) analyzeWhileStmt(ws *parser.WhileStmt) *WhileStmt {
	return &WhileStmt{
		Condition: a.analyzeExpr(ws.Condition),
		Body:      a.analyzeBlock(ws.Body),
	}
}
