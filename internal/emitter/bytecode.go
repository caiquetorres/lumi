package emitter

const (
	LoadConst byte = iota + 1
	End
	BeginScope
	EndScope

	FnDecl
	VarDecl

	GetSymbol
	Call
	Return

	Pop
)
