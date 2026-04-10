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

	Pub
	Fun
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
	case Pub:
		return "pub"
	case Fun:
		return "fun"
	default:
		return "unknown"
	}
}
