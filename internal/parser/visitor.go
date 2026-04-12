package parser

type Visitor interface {
	VisitAst(*Ast) error
	VisitFunDecl(*FunDecl) error
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
			return w.walkFunDecl(v, s)
		}
	}

	return nil
}

func (w *walker) walkFunDecl(v Visitor, fd *FunDecl) error {
	return v.VisitFunDecl(fd)
}
