package emitter

const (
	PushTrue byte = iota + 1
	PushFalse
	PushInt
	PushString
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

	StoreLocal
	LoadLocal

	Call
	Return

	JumpTo
	JumpIfFalse

	Pop
)
