package lexer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/caiquetorres/lumi/internal/span"
	"github.com/caiquetorres/lumi/internal/token"
)

type Lexer struct {
	start, end int

	currToken token.Token
	currErr   error

	b *bufio.Reader

	currLexeme string
	symTable   *symbolTable
}

func New(r io.Reader) *Lexer {
	l := Lexer{
		b:        bufio.NewReader(r), // the default buffer size is 4096
		symTable: newSymbolTable(),
	}

	l.currToken, l.currErr = l.next()

	return &l
}

func (l *Lexer) Peek() (token.Token, error) {
	return l.currToken, l.currErr
}

func (l *Lexer) Next() (token.Token, error) {
	currToken, currErr := l.currToken, l.currErr

	l.currToken, l.currErr = l.next()

	return currToken, currErr
}

func (l *Lexer) Lexeme(tok token.Token) string {
	return l.symTable.lookup(tok.SymbolID())
}

func (l *Lexer) DebugTable() {
	for sym, id := range l.symTable.index {
		fmt.Printf("%d: %s\n", id, sym)
	}
}

func (l *Lexer) nextRune() (rune, error) {
	r, _, err := l.b.ReadRune()

	l.extendSpan()
	l.currLexeme += string(r)

	return r, err
}

func (l *Lexer) peekRune() (rune, error) {
	r, _, err := l.b.ReadRune()
	if err != nil {
		return 0, err
	}

	if err := l.b.UnreadRune(); err != nil {
		return 0, err
	}

	return r, nil
}

func (l *Lexer) isAtEOF() bool {
	_, err := l.peekRune()
	return err == io.EOF
}

func (l *Lexer) newToken(k token.Kind) token.Token {
	id := l.symTable.intern(l.currLexeme)

	tok := token.NewWithSymbol(id, k, span.New(l.start, l.end))
	l.resetSpan()
	return tok
}

func (l *Lexer) resetSpan() {
	l.start = l.end
}

func (l *Lexer) extendSpan() {
	l.end++
}
