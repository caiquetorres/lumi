package parser

type Visitor interface {
	BeforeAst(*Ast)
	BeforeFunDecl(*FunDecl)
	AfterFunDecl(*FunDecl)
	BeforeVarDecl(*VarDecl)
	AfterVarDecl(*VarDecl)
	BeforeParam(*Param)
	AfterParam(*Param)
	BeforeLiteralExpr(*LiteralExpr)
	BeforeIdentifierExpr(*IdentifierExpr)

	BeforeCallExpr(*CallExpr)
	AfterCallExpr(*CallExpr)

	BeforeBlockExpr(*Block)
	AfterBlockExpr(*Block)

	BeforeBreakStmt(*BreakStmt)
	AfterBreakStmt(*BreakStmt)

	BeforeContinueStmt(*ContinueStmt)
	AfterContinueStmt(*ContinueStmt)

	AfterIfCondition(*IfStmt)
	AfterIfThenBlock(*IfStmt)
	AfterElseBlock(*IfStmt)

	BeforeWhileCondition(*WhileStmt)
	AfterWhileCondition(*WhileStmt)
	AfterWhileBody(*WhileStmt)

	BeforeReturnStmt(*ReturnStmt)
	AfterReturnStmt(*ReturnStmt)

	AfterStmt(Stmt)
}
