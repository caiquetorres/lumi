package emitter

const (
	LoadConst byte = iota + 1

	PushInt
	PushTrue
	PushFalse
	PushNativeFn
	PushFn

	FnDecl

	Not

	Add
	Sub
	Mul
	Div
	Eq
	Less
	LessEq

	// GetSymbol
	SetSymbol
	DefineSymbol

	Call
	Return

	JumpTo
	JumpIfFalse

	Pop
)
