package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type IfStmt struct {
	Condition Expr
	Then      *Block
	Else      *Block
}

func ifStmt(is *parser.IfStmt) *IfStmt {
	return &IfStmt{
		Condition: exprN(is.Condition),
		Then:      block(is.Then),
		Else:      block(is.Else),
	}
}
