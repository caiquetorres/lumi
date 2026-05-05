package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/caiquetorres/lumi/internal/span"
)

var (
	ErrUnexpectedToken = errors.New("unexpected token")
	ErrUnexpectedEOF   = errors.New("unexpected end of file")
)

type ParseError struct {
	Err  error
	Span span.Span
}

func (p *ParseError) Error() string {
	return p.Err.Error()
}

var _ error = (*ParseError)(nil)

func Format(err error, r io.ReadSeeker) string {
	parseErr, ok := err.(*ParseError)
	if !ok {
		return err.Error()
	}

	r.Seek(0, io.SeekStart)
	contents, _ := io.ReadAll(r)

	start := int(parseErr.Span.Start())

	lineStart := start
	for lineStart > 0 && contents[lineStart-1] != '\n' {
		lineStart--
	}

	lineEnd := start
	for lineEnd < len(contents) && contents[lineEnd] != '\n' {
		lineEnd++
	}

	lineNum := 1
	for i := range start {
		if contents[i] == '\n' {
			lineNum++
		}
	}

	prevLineEnd := lineStart - 1 // index of the '\n' before current line, or -1
	prevLineStart := -1
	if prevLineEnd >= 0 {
		prevLineStart = prevLineEnd
		for prevLineStart > 0 && contents[prevLineStart-1] != '\n' {
			prevLineStart--
		}
	}

	nextLineStart, nextLineEnd := -1, -1
	if lineEnd < len(contents) {
		nextLineStart = lineEnd + 1
		nextLineEnd = nextLineStart
		for nextLineEnd < len(contents) && contents[nextLineEnd] != '\n' {
			nextLineEnd++
		}
	}

	col := start - lineStart
	line := string(contents[lineStart:lineEnd])

	caretLen := int(parseErr.Span.End()) - start
	if start+caretLen > lineEnd {
		caretLen = lineEnd - start
	}
	if caretLen < 1 {
		caretLen = 1
	}

	maxLineNum := lineNum
	if nextLineStart >= 0 {
		maxLineNum = lineNum + 1
	}
	width := len(fmt.Sprintf("%d", maxLineNum))
	pad := strings.Repeat(" ", width)

	numFmt := func(n int) string {
		s := fmt.Sprintf("%d", n)
		return strings.Repeat(" ", width-len(s)) + s
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "error: %s\n", parseErr.Err.Error())
	fmt.Fprintf(&sb, " %s --> line %d:%d\n", pad, lineNum, col+1)
	fmt.Fprintf(&sb, " %s |\n", pad)

	if prevLineStart >= 0 {
		fmt.Fprintf(&sb, " %s | %s\n", numFmt(lineNum-1), string(contents[prevLineStart:prevLineEnd]))
	}

	fmt.Fprintf(&sb, " %s | %s\n", numFmt(lineNum), line)
	fmt.Fprintf(&sb, " %s | %s%s\n", pad, strings.Repeat(" ", col), strings.Repeat("^", caretLen))

	if nextLineStart >= 0 {
		fmt.Fprintf(&sb, " %s | %s\n", numFmt(lineNum+1), string(contents[nextLineStart:nextLineEnd]))
	}

	fmt.Fprintf(&sb, " %s |\n", pad)

	return sb.String()
}
