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

	BeforeBlockExpr(*Block) error
	AfterBlockExpr(*Block) error

	BeforeBreakStmt(*Break) error
	AfterBreakStmt(*Break) error

	BeforeContinueStmt(*Continue) error
	AfterContinueStmt(*Continue) error

	AfterIfCondition(*If) error
	AfterIfThenBlock(*If) error
	AfterElseBlock(*If) error

	BeforeWhileCondition(*While) error
	AfterWhileCondition(*While) error
	AfterWhileBody(*While) error

	BeforeReturnStmt(*Return) error
	AfterReturnStmt(*Return) error

	AfterStmt(Stmt) error
}
