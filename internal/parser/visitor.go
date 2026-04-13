package parser

type Visitor interface {
	BeforeAst(*Ast) error
	BeforeFunDecl(*FunDecl) error
	AfterFunDecl(*FunDecl) error
	BeforeLiteralExpr(*LiteralExpr) error
	BeforeIdentifierExpr(*IdentifierExpr) error
	AfterCallExpr(*CallExpr) error
	AfterStmt(Stmt) error
}
