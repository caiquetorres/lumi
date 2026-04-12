package parser

type Visitor interface {
	VisitAst(*Ast) error
	VisitFunDeclStart(*FunDecl) error
	VisitFunDeclEnd(*FunDecl) error
	VisitLiteralExpr(*LiteralExpr) error
	VisitExprEnd(Expression) error
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
		switch e := expr.(type) {
		case *LiteralExpr:
			if err := v.VisitLiteralExpr(e); err != nil {
				return err
			}
		}
		if err := v.VisitExprEnd(expr); err != nil {
			return err
		}
	}

	if err := v.VisitFunDeclEnd(fd); err != nil {
		return err
	}

	return nil
}
