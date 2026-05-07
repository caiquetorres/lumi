package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type TopLevelStmt any

type Ast struct {
	Statements []TopLevelStmt
	Node       *parser.Ast
}

type FunDecl struct {
	Node *parser.FunDecl
}

func (w *Analyzer) WalkAst(ast *parser.Ast) *Ast {
	stmts := make([]TopLevelStmt, len(ast.Statements))
	for _, stmt := range ast.Statements {
		var aStmt TopLevelStmt
		switch s := stmt.(type) {
		case *parser.FunDecl:
			aStmt = w.walkFunDecl(s)
		}

		stmts = append(stmts, aStmt)
	}

	return &Ast{
		Statements: stmts,
		Node:       ast,
	}
}

func (w *Analyzer) walkFunDecl(fd *parser.FunDecl) *FunDecl {
	return &FunDecl{
		Node: fd,
	}
}
