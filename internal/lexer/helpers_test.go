package lexer

import (
	"bufio"
	"io"
)

func newRaw(r io.Reader) *Lexer {
	return &Lexer{b: bufio.NewReader(r)}
}
