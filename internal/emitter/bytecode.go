package emitter

const (
	LoadConst byte = iota + 1
	BeginScope
	EndScope

	FnDecl

	GetSymbol
	DefineSymbol

	Call
	Return

	JumpTo
	JumpIfFalse

	Pop
)
