package parser

func Walk(v Visitor, ast *Ast) {
	w := &walker{}
	w.walkAst(v, ast)
}

type walker struct{}

func (w *walker) walkAst(v Visitor, ast *Ast) {
	v.BeforeAst(ast)

	for _, stmt := range ast.Statements {
		switch s := stmt.(type) {
		case *FunDecl:
			w.walkFunDecl(v, s)
		}
	}
}

func (w *walker) walkVarDecl(v Visitor, vd *VarDecl) {
	v.BeforeVarDecl(vd)

	w.walkExpr(v, vd.Expr)

	v.AfterVarDecl(vd)
}

func (w *walker) walkFunDecl(v Visitor, fd *FunDecl) {
	v.BeforeFunDecl(fd)

	for _, param := range fd.Params {
		// REVIEW: We definitely don't need two separate methods (before and after) here
		v.BeforeParam(&param)
		v.AfterParam(&param)
	}

	for _, expr := range fd.Body {
		w.walkStmt(v, expr)
	}

	v.AfterFunDecl(fd)
}

func (w *walker) walkStmt(v Visitor, stmt Stmt) {
	switch s := stmt.(type) {
	case *VarDecl:
		w.walkVarDecl(v, s)
	case *ReturnStmt:
		w.walkReturn(v, s)
	case *IfStmt:
		w.walkIf(v, s)
	case *WhileStmt:
		w.walkWhile(v, s)
	case *BreakStmt:
		w.walkBreak(v, s)
	case *ContinueStmt:
		w.walkContinue(v, s)
	case *Block:
		w.walkBlockStmt(v, s)
	default:
		w.walkExpr(v, stmt.(Expr))
	}

	v.AfterStmt(stmt)
}

func (w *walker) walkWhile(v Visitor, whileStmt *WhileStmt) {
	v.BeforeWhileCondition(whileStmt)

	w.walkExpr(v, whileStmt.Condition)

	v.AfterWhileCondition(whileStmt)

	w.walkBlockStmt(v, whileStmt.Body)

	v.AfterWhileBody(whileStmt)
}

func (w *walker) walkIf(v Visitor, ifStmt *IfStmt) error {
	w.walkExpr(v, ifStmt.Condition)

	v.AfterIfCondition(ifStmt)

	w.walkBlockStmt(v, ifStmt.Then)

	v.AfterIfThenBlock(ifStmt)

	if ifStmt.Else != nil {
		w.walkBlockStmt(v, ifStmt.Else)

		v.AfterElseBlock(ifStmt)
	}

	return nil
}

func (w *walker) walkReturn(v Visitor, r *ReturnStmt) {
	v.BeforeReturnStmt(r)

	if r.Expr != nil {
		w.walkExpr(v, r.Expr)
	}

	v.AfterReturnStmt(r)
}

func (w *walker) walkBreak(v Visitor, b *BreakStmt) {
	// REVIEW: We definitely don't need two separate methods (before and after) here
	v.BeforeBreakStmt(b)
	v.AfterBreakStmt(b)
}

func (w *walker) walkContinue(v Visitor, c *ContinueStmt) {
	// REVIEW: We definitely don't need two separate methods (before and after) here
	v.BeforeContinueStmt(c)
	v.AfterContinueStmt(c)
}

func (w *walker) walkExpr(v Visitor, expr Expr) {
	switch e := expr.(type) {
	case *IdentifierExpr:
		w.walkIdentifierExpr(v, e)
	case *LiteralExpr:
		w.walkLiteralExpr(v, e)
	case *CallExpr:
		w.walkCallExpr(v, e)
	}
}

func (w *walker) walkIdentifierExpr(v Visitor, ie *IdentifierExpr) {
	v.BeforeIdentifierExpr(ie)
}

func (w *walker) walkLiteralExpr(v Visitor, le *LiteralExpr) {
	v.BeforeLiteralExpr(le)
}

func (w *walker) walkCallExpr(v Visitor, ce *CallExpr) {
	v.BeforeCallExpr(ce)

	for _, arg := range ce.Args {
		w.walkExpr(v, arg)
	}

	w.walkExpr(v, ce.Callee)

	v.AfterCallExpr(ce)
}

func (w *walker) walkBlockStmt(v Visitor, be *Block) {
	v.BeforeBlockExpr(be)

	for _, stmt := range be.Stmts {
		w.walkStmt(v, stmt)
	}

	v.AfterBlockExpr(be)
}
