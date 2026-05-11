package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Stmt any

func (t *TypeChecker) analyzeStmt(s parser.Stmt) Stmt {
	switch n := s.(type) {
	case *parser.VarDecl:
		return t.analyzeVarDecl(n)
	case *parser.IfStmt:
		return t.analyzeIfStmt(n)
	case *parser.ReturnStmt:
		return t.analyzeReturnStmt(n)
	case *parser.ForStmt:
		return t.analyzeForStmt(n)
	case *parser.WhileStmt:
		return t.analyzeWhileStmt(n)
	case *parser.BreakStmt:
		return t.analyzeBreakStmt(n)
	case *parser.ContinueStmt:
		return t.analyzeContinueStmt(n)
	case *parser.Loop:
		return t.analyzeLoop(n)
	case *parser.Block:
		return t.analyzeBlock(n)
	default:
		return t.analyzeExpr(s.(parser.Expr))
	}
}

type Block struct {
	Stmts []Stmt
}

func (t *TypeChecker) analyzeBlock(b *parser.Block) *Block {
	if b == nil {
		return nil
	}

	stmts := make([]Stmt, len(b.Stmts))
	for i, s := range b.Stmts {
		stmts[i] = t.analyzeStmt(s)
	}

	return &Block{
		Stmts: stmts,
	}
}
