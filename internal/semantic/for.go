package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type ForStmt struct {
	Init Stmt
	Cond Expr
	Inc  Stmt
	Body *Block
}

func (t *TypeChecker) analyzeForStmt(fs *parser.ForStmt) *ForStmt {
	return &ForStmt{
		Init: t.analyzeStmt(fs.Init),
		Cond: t.analyzeExpr(fs.Cond),
		Inc:  t.analyzeStmt(fs.Inc),
		Body: t.analyzeBlock(fs.Body),
	}
}
