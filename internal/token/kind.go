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
	Equals

	String

	Pub
	Fun
	Let
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
	case String:
		return "string"
	case Pub:
		return "pub"
	case Fun:
		return "fun"
	case Let:
		return "let"
	default:
		return "unknown"
	}
}
