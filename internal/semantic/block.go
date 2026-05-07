package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Stmt any

func stmtN(s parser.Stmt) Stmt {
	switch n := s.(type) {
	case *parser.VarDecl:
		return varDecl(n)
	case *parser.IfStmt:
		return ifStmt(n)
	case *parser.ReturnStmt:
		return returnStmt(n)
	case *parser.ForStmt:
		return forStmt(n)
	case *parser.WhileStmt:
		return whileStmt(n)
	case *parser.BreakStmt:
		return breakStmt(n)
	case *parser.ContinueStmt:
		return continueStmt(n)
	case *parser.Loop:
		return loop(n)
	case *parser.Block:
		return block(n)
	default:
		return exprN(s.(parser.Expr))
	}
}

type Block struct {
	Stmts []Stmt
}

func block(b *parser.Block) *Block {
	if b == nil {
		return nil
	}
	stmts := make([]Stmt, len(b.Stmts))
	for i, s := range b.Stmts {
		stmts[i] = stmtN(s)
	}
	return &Block{
		Stmts: stmts,
	}
}
