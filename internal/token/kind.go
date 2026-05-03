package token

type Kind int

const (
	Bad Kind = iota
	EOF

	Identifier

	OpenParen
	CloseParen
	OpenBrace
	CloseBrace
	Semicolon
	Comma
	NewLine

	Plus
	PlusEqual

	Minus
	MinusEqual

	Star
	StarEqual

	Slash
	SlashEqual

	Dot
	DotDot
	DotDotEqual

	Equal
	EqualEqual
	Bang
	BangEqual
	Less
	LessEqual
	Greater
	GreaterEqual

	String
	Int

	False
	True

	Pub
	Fun
	Let
	Return
	Break
	Continue
	If
	Else
	While
	For
	In
)

func (k Kind) String() string {
	switch k {
	case Bad:
		return "bad"
	case EOF:
		return "eof"
	case Identifier:
		return "identifier"
	case OpenParen:
		return "("
	case CloseParen:
		return ")"
	case OpenBrace:
		return "{"
	case CloseBrace:
		return "}"
	case Semicolon:
		return ";"
	case Comma:
		return ","
	case False:
		return "false"
	case True:
		return "true"
	case Plus:
		return "+"
	case PlusEqual:
		return "+="
	case Minus:
		return "-"
	case MinusEqual:
		return "-="
	case Star:
		return "*"
	case StarEqual:
		return "*="
	case Slash:
		return "/"
	case SlashEqual:
		return "/="
	case Equal:
		return "="
	case EqualEqual:
		return "=="
	case Bang:
		return "!"
	case BangEqual:
		return "!="
	case Less:
		return "<"
	case LessEqual:
		return "<="
	case Greater:
		return ">"
	case GreaterEqual:
		return ">="
	case NewLine:
		return "new line"
	case String:
		return "string"
	case Int:
		return "number"
	case Pub:
		return "pub"
	case Fun:
		return "fun"
	case Let:
		return "let"
	case Return:
		return "return"
	case Break:
		return "break"
	case Continue:
		return "continue"
	case If:
		return "if"
	case Else:
		return "else"
	case While:
		return "while"
	default:
		return "unknown"
	}
}
