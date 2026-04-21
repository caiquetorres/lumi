package parser

func Walk(v Visitor, ast *Ast) error {
	w := &walker{}
	return w.walkAst(v, ast)
}

type walker struct{}

func (w *walker) walkAst(v Visitor, ast *Ast) error {
	if err := v.BeforeAst(ast); err != nil {
		return err
	}

	for _, stmt := range ast.Statements {
		var err error

		switch s := stmt.(type) {
		case *FunDecl:
			err = w.walkFunDecl(v, s)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *walker) walkVarDecl(v Visitor, vd *VarDecl) error {
	if err := v.BeforeVarDecl(vd); err != nil {
		return err
	}

	if err := w.walkExpr(v, vd.Expr); err != nil {
		return err
	}

	return v.AfterVarDecl(vd)
}

func (w *walker) walkFunDecl(v Visitor, fd *FunDecl) error {
	if err := v.BeforeFunDecl(fd); err != nil {
		return err
	}

	for _, param := range fd.Params {
		if err := v.BeforeParam(&param); err != nil {
			return err
		}

		if err := v.AfterParam(&param); err != nil {
			return err
		}
	}

	for _, expr := range fd.Body {
		if err := w.walkStmt(v, expr); err != nil {
			return err
		}
	}

	return v.AfterFunDecl(fd)
}

func (w *walker) walkStmt(v Visitor, stmt Stmt) error {
	var err error

	switch s := stmt.(type) {
	case *VarDecl:
		err = w.walkVarDecl(v, s)
	case *Return:
		err = w.walkReturn(v, s)
	case *Break:
		err = w.walkBreak(v, s)
	default:
		err = w.walkExpr(v, stmt.(Expr))
	}

	if err != nil {
		return err
	}

	return v.AfterStmt(stmt)
}

func (w *walker) walkReturn(v Visitor, r *Return) error {
	if err := v.BeforeReturnStmt(r); err != nil {
		return err
	}

	if r.Expr != nil {
		if err := w.walkExpr(v, r.Expr); err != nil {
			return err
		}
	}

	return v.AfterReturnStmt(r)
}

func (w *walker) walkBreak(v Visitor, b *Break) error {
	if err := v.BeforeBreakStmt(b); err != nil {
		return err
	}

	return v.AfterBreakStmt(b)
}

func (w *walker) walkExpr(v Visitor, expr Expr) error {
	var err error

	switch e := expr.(type) {
	case *IdentifierExpr:
		err = w.walkIdentifierExpr(v, e)
	case *LiteralExpr:
		err = w.walkLiteralExpr(v, e)
	case *CallExpr:
		err = w.walkCallExpr(v, e)
	case *BlockExpr:
		err = w.walkBlockExpr(v, e)
	}

	return err
}

func (w *walker) walkIdentifierExpr(v Visitor, ie *IdentifierExpr) error {
	return v.BeforeIdentifierExpr(ie)
}

func (w *walker) walkLiteralExpr(v Visitor, le *LiteralExpr) error {
	return v.BeforeLiteralExpr(le)
}

func (w *walker) walkCallExpr(v Visitor, ce *CallExpr) error {
	if err := v.BeforeCallExpr(ce); err != nil {
		return err
	}

	for _, arg := range ce.Args {
		if err := w.walkExpr(v, arg); err != nil {
			return err
		}
	}

	if err := w.walkExpr(v, ce.Callee); err != nil {
		return err
	}

	return v.AfterCallExpr(ce)
}

func (w *walker) walkBlockExpr(v Visitor, be *BlockExpr) error {
	if err := v.BeforeBlockExpr(be); err != nil {
		return err
	}

	for _, stmt := range be.Stmts {
		if err := w.walkStmt(v, stmt); err != nil {
			return err
		}
	}

	return v.AfterBlockExpr(be)
}
