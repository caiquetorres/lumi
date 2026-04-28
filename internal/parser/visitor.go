package parser

type Visitor interface {
	BeforeAst(*Ast) error
	BeforeFunDecl(*FunDecl) error
	AfterFunDecl(*FunDecl) error
	BeforeVarDecl(*VarDecl) error
	AfterVarDecl(*VarDecl) error
	BeforeParam(*Param) error
	AfterParam(*Param) error
	BeforeLiteralExpr(*LiteralExpr) error
	BeforeIdentifierExpr(*IdentifierExpr) error

	BeforeCallExpr(*CallExpr) error
	AfterCallExpr(*CallExpr) error

	BeforeBlockExpr(*BlockExpr) error
	AfterBlockExpr(*BlockExpr) error

	BeforeBreakStmt(*Break) error
	AfterBreakStmt(*Break) error

	AfterIfCondition(*If) error
	AfterIfThenBlock(*If) error

	BeforeReturnStmt(*Return) error
	AfterReturnStmt(*Return) error

	AfterStmt(Stmt) error
}
