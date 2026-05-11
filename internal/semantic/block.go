package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Stmt any

func (a *TypeChecker) analyzeStmt(s parser.Stmt) Stmt {
	switch n := s.(type) {
	case *parser.VarDecl:
		return a.analyzeVarDecl(n)
	case *parser.IfStmt:
		return a.analyzeIfStmt(n)
	case *parser.ReturnStmt:
		return a.analyzeReturnStmt(n)
	case *parser.ForStmt:
		return a.analyzeForStmt(n)
	case *parser.WhileStmt:
		return a.analyzeWhileStmt(n)
	case *parser.BreakStmt:
		return a.analyzeBreakStmt(n)
	case *parser.ContinueStmt:
		return a.analyzeContinueStmt(n)
	case *parser.Loop:
		return a.analyzeLoop(n)
	case *parser.Block:
		return a.analyzeBlock(n)
	default:
		return a.analyzeExpr(s.(parser.Expr))
	}
}

type Block struct {
	Stmts []Stmt
}

func (a *TypeChecker) analyzeBlock(b *parser.Block) *Block {
	if b == nil {
		return nil
	}

	stmts := make([]Stmt, len(b.Stmts))
	for i, s := range b.Stmts {
		stmts[i] = a.analyzeStmt(s)
	}

	return &Block{
		Stmts: stmts,
	}
}
