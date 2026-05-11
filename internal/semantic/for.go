package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ForStmt struct {
	Init Stmt
	Cond Expr
	Inc  Stmt
	Body *Block
}

func (a *TypeChecker) analyzeForStmt(fs *parser.ForStmt) *ForStmt {
	return &ForStmt{
		Init: a.analyzeStmt(fs.Init),
		Cond: a.analyzeExpr(fs.Cond),
		Inc:  a.analyzeStmt(fs.Inc),
		Body: a.analyzeBlock(fs.Body),
	}
}
