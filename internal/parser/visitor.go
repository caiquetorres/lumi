package parser

type Visitor interface {
	BeforeAst(*Ast)
	BeforeFunDecl(*FunDecl)
	AfterFunDecl(*FunDecl)

	BeforeVarDecl(*Let)
	AfterVarDecl(*Let)

	BeforeAssignment(*Binding)
	AfterAssignment(*Binding)

	BeforeParam(*Param)
	AfterParam(*Param)
	BeforeLiteralExpr(*LiteralExpr)
	BeforeIdentifierExpr(*IdentifierExpr)

	BeforeCallExpr(*CallExpr)
	AfterCallExpr(*CallExpr)

	BeforeBinaryExpr(*BinaryExpr)
	AfterBinaryExpr(*BinaryExpr)

	BeforeBlockStmt(*Block)
	AfterBlockStmt(*Block)

	BeforeBreakStmt(*Break)
	AfterBreakStmt(*Break)

	BeforeContinueStmt(*Continue)
	AfterContinueStmt(*Continue)

	AfterIfCondition(*If)
	AfterIfThenBlock(*If)
	AfterElseBlock(*If)

	BeforeLoopBody(*Loop)
	AfterLoopBody(*Loop)

	BeforeWhileCondition(*WhileStmt)
	AfterWhileCondition(*WhileStmt)
	AfterWhileBody(*WhileStmt)

	BeforeForInit(*For)
	AfterForInit(*For)
	BeforeForInc(*For)
	AfterForInc(*For)
	BeforeForCond(*For)
	AfterForCond(*For)
	AfterForBody(*For)

	BeforeReturnStmt(*ReturnStmt)
	AfterReturnStmt(*ReturnStmt)

	AfterStmt(Stmt)
}
