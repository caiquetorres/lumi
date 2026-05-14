package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/caiquetorres/lumi/internal/emitter/v2"
	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/semantic"
	"github.com/caiquetorres/lumi/internal/vm/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompilation(t *testing.T) {
	files, err := os.ReadDir("tests")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".lumi") {
			continue
		}

		name := strings.TrimSuffix(f.Name(), ".lumi")
		srcPath := filepath.Join("tests", name+".lumi")
		expPath := filepath.Join("tests", name+".expected")

		t.Run(name, func(t *testing.T) {
			actual, err := compileToBytecodeString(srcPath)
			require.NoError(t, err)

			expected, err := os.ReadFile(expPath)
			require.NoError(t, err)

			assert.Equal(t, normalizeLines(string(expected)), normalizeLines(actual))
		})
	}
}

const printlnSrc = `fun main() {
	let s = 1 + 2
	println(s)
}`

func TestExec(t *testing.T) {
	t.Run("println", func(t *testing.T) {
		r := compileToLumi(t, printlnSrc)

		var out bytes.Buffer
		if err := vm.ExecWithWriter(r, &out); err != nil {
			t.Fatalf("exec failed: %v", err)
		}

		got := out.String()
		want := "3\n"

		assert.Equal(t, want, got)
	})
}

func compileToBytecodeString(srcPath string) (string, error) {
	f, err := os.Open(srcPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	lex := lexer.New(f)
	p := parser.New(lex)

	ast, err := p.Parse()
	if err != nil {
		return "", err
	}

	sAst, err := semantic.Analyze(ast, lex)
	if err != nil {
		return "", err
	}

	ch, err := emitter.Emit(sAst, lex)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	emitter.NewDisassembler(&b, ch).Disassemble()

	return b.String(), nil
}

func normalizeLines(s string) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t")
	}
	return strings.Join(lines, "\n")
}

func compileToLumi(t *testing.T, src string) *bytes.Reader {
	t.Helper()

	lex := lexer.New(strings.NewReader(src))
	par := parser.New(lex)

	ast, err := par.Parse()
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	sAst, err := semantic.Analyze(ast, lex)
	if err != nil {
		t.Fatalf("semantic analyze failed: %v", err)
	}

	ch, err := emitter.Emit(sAst, lex)
	if err != nil {
		t.Fatalf("emit failed: %v", err)
	}

	var buf bytes.Buffer
	if err := emitter.WriteLumiFile(ch, &buf); err != nil {
		t.Fatalf("writing lumi file failed: %v", err)
	}

	return bytes.NewReader(buf.Bytes())
}
