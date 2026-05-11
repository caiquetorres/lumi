package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type IfStmt struct {
	Condition Expr
	Then      *Block
	Else      *Block
}

func (a *Analyzer) analyzeIfStmt(is *parser.IfStmt) *IfStmt {
	return &IfStmt{
		Condition: a.analyzeExpr(is.Condition),
		Then:      a.analyzeBlock(is.Then),
		Else:      a.analyzeBlock(is.Else),
	}
}
