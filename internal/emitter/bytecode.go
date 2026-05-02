package emitter

const (
	LoadConst byte = iota + 1

	PushString
	PushInt
	PushTrue
	PushFalse
	PushNativeFn
	PushFn

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

	StoreLocal
	LoadLocal

	Call
	Return

	JumpTo
	JumpIfFalse

	Pop
)
