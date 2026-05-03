package parser

type Visitor interface {
	BeforeAst(*Ast)
	BeforeFunDecl(*FunDecl)
	AfterFunDecl(*FunDecl)

	BeforeVarDecl(*VarDecl)
	AfterVarDecl(*VarDecl)

	BeforeAssignment(*Assignment)
	AfterAssignment(*Assignment)

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
