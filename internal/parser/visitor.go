package parser

type Visitor interface {
	BeforeAst(*Ast) error
	BeforeFunDecl(*FunDecl) error
	AfterFunDecl(*FunDecl) error
	BeforeParam(*Param) error
	AfterParam(*Param) error
	BeforeLiteralExpr(*LiteralExpr) error
	BeforeIdentifierExpr(*IdentifierExpr) error
	BeforeCallExpr(*CallExpr) error
	AfterCallExpr(*CallExpr) error
	AfterStmt(Stmt) error
}
