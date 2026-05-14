package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type For struct {
	Init Stmt
	Cond Expr
	Inc  Stmt
	Body *Block
}

func (t *TypeChecker) analyzeForStmt(fs *parser.For) *For {
	return &For{
		Init: t.analyzeStmt(fs.Init),
		Cond: t.analyzeExpr(fs.Cond),
		Inc:  t.analyzeStmt(fs.Inc),
		Body: t.analyzeBlock(fs.Body),
	}
}
