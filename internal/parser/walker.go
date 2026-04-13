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

func (w *walker) walkFunDecl(v Visitor, fd *FunDecl) error {
	if err := v.BeforeFunDecl(fd); err != nil {
		return err
	}

	for _, expr := range fd.Body {
		if err := w.walkStmt(v, expr); err != nil {
			return err
		}
	}

	return v.AfterFunDecl(fd)
}

func (w *walker) walkStmt(v Visitor, stmt Stmt) error {
	if err := w.walkExpr(v, stmt.(Expr)); err != nil {
		return err
	}

	return v.AfterStmt(stmt)
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
	if err := w.walkExpr(v, ce.Callee); err != nil {
		return err
	}

	// TODO: walk arguments

	return v.AfterCallExpr(ce)
}
