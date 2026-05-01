package emitter

const (
	LoadConst byte = iota + 1
	BeginScope
	EndScope

	FnDecl

	Add
	Sub
	Mul
	Div

	GetSymbol
	DefineSymbol

	Call
	Return

	JumpTo
	JumpIfFalse

	Pop
)
