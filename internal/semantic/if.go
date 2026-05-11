package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type IfStmt struct {
	Condition Expr
	Then      *Block
	Else      *Block
}

func (t *TypeChecker) analyzeIfStmt(is *parser.IfStmt) *IfStmt {
	return &IfStmt{
		Condition: t.analyzeExpr(is.Condition),
		Then:      t.analyzeBlock(is.Then),
		Else:      t.analyzeBlock(is.Else),
	}
}
