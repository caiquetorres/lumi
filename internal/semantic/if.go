package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type If struct {
	Condition Expr
	Then      *Block
	Else      *Block
}

func (t *TypeChecker) analyzeIfStmt(is *parser.If) *If {
	return &If{
		Condition: t.analyzeExpr(is.Condition),
		Then:      t.analyzeBlock(is.Then),
		Else:      t.analyzeBlock(is.Else),
	}
}
