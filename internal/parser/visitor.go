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

	BeforeBreakStmt(*Break)
	AfterBreakStmt(*Break)

	BeforeContinueStmt(*Continue)
	AfterContinueStmt(*Continue)

	AfterIfCondition(*If)
	AfterIfThenBlock(*If)
	AfterElseBlock(*If)

	BeforeWhileCondition(*While)
	AfterWhileCondition(*While)
	AfterWhileBody(*While)

	BeforeReturnStmt(*Return)
	AfterReturnStmt(*Return)

	AfterStmt(Stmt)
}
