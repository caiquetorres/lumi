package emitter

const (
	LoadConst byte = iota + 1
	End
	BeginScope
	EndScope

	FnDecl

	GetSymbol
	DefineSymbol

	Call
	Return

	Jump
	Pop
)
