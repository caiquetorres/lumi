package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ForStmt struct {
	Init Stmt
	Cond Expr
	Inc  Stmt
	Body *Block
}

func forStmt(fs *parser.ForStmt) *ForStmt {
	return &ForStmt{
		Init: stmtN(fs.Init),
		Cond: exprN(fs.Cond),
		Inc:  stmtN(fs.Inc),
		Body: block(fs.Body),
	}
}
