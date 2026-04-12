package emitter

const (
	LoadConst byte = iota + 1
	End
	BeginScope
	EndScope
	DeclFun

	Pop
)
