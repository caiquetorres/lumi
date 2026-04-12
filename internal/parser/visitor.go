package parser

type Visitor interface {
	VisitAst(*Ast) error
	VisitFunDeclStart(*FunDecl) error
	VisitFunDeclEnd(*FunDecl) error

	VisitLiteralExpr(*LiteralExpr) error
	VisitIdentifierExpr(*IdentifierExpr) error
	VisitCallExpr(*CallExpr) error

	VisitStmtEnd(Expr) error
}

func Walk(v Visitor, ast *Ast) error {
	w := &walker{}
	return w.walkAsk(v, ast)
}

type walker struct{}

func (w *walker) walkAsk(v Visitor, ast *Ast) error {
	if err := v.VisitAst(ast); err != nil {
		return err
	}

	for _, stmt := range ast.Statements {
		switch s := stmt.(type) {
		case *FunDecl:
			if err := w.walkFunDecl(v, s); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *walker) walkFunDecl(v Visitor, fd *FunDecl) error {
	if err := v.VisitFunDeclStart(fd); err != nil {
		return err
	}

	for _, expr := range fd.Body {
		if err := w.walkExpr(v, expr); err != nil {
			return err
		}

		if err := v.VisitStmtEnd(expr); err != nil {
			return err
		}
	}

	if err := v.VisitFunDeclEnd(fd); err != nil {
		return err
	}

	return nil
}

func (w *walker) walkExpr(v Visitor, expr Expr) error {
	switch e := expr.(type) {
	case *IdentifierExpr:
		if err := v.VisitIdentifierExpr(e); err != nil {
			return err
		}
	case *LiteralExpr:
		if err := v.VisitLiteralExpr(e); err != nil {
			return err
		}
	case *CallExpr:
		if err := w.walkCallExpr(v, e); err != nil {
			return err
		}
	}

	return nil
}

func (w *walker) walkCallExpr(v Visitor, ce *CallExpr) error {
	if err := w.walkExpr(v, ce.Callee); err != nil {
		return err
	}

	// TODO: walk arguments

	return v.VisitCallExpr(ce)
}
