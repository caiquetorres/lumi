package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"

	"github.com/caiquetorres/lumi/internal/emitter"
	"github.com/caiquetorres/lumi/internal/lexer"
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/semantic"
	"github.com/caiquetorres/lumi/internal/vm/v2"
)

func main() {
	mode := "build"
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "build", "run":
			mode = os.Args[1]
			os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
		}
	}

	switch mode {
	case "run":
		args := parseRunArgs()

		f, err := os.Open(args.filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := vm.Exec(f); err != nil {
			log.Fatal(err)
		}

	case "build":
		args := parseArgs()
		log.Printf("running with file: %s", args.filePath)

		f, err := os.Open(args.filePath)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Remove(args.outPath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}

		o, err := os.Create(args.outPath)
		if err != nil {
			log.Fatal(err)
		}

		if err := compilePipeline(f, o, args.disassemble); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", mode)
	}
}

func compilePipeline(src io.Reader, out io.Writer, disassemble bool) error {
	var (
		l = lexer.New(src)
		p = parser.New(l)
	)

	ast, err := p.Parse()
	if err != nil {
		return err
	}

	if err := semantic.Analyze(ast); err != nil {
		return err
	}

	ch, err := emitter.Emit(ast, l, out)
	if err != nil {
		return err
	}

	if disassemble {
		emitter.
			NewDisassembler(os.Stdout, ch).
			Disassemble()
	}

	return nil
}

type arguments struct {
	filePath    string
	outPath     string
	disassemble bool
}

type runArguments struct {
	filePath string
}

func parseArgs() arguments {
	file := flag.String("file", "", "path to source file")
	out := flag.String("out", "", "path to output file")
	disassemble := flag.Bool("disassemble", false, "enable disassemble mode")

	flag.Parse()

	if file == nil || *file == "" {
		log.Fatal("file argument is required")
	}

	if out == nil || *out == "" {
		log.Fatal("out argument is required")
	}

	if disassemble == nil {
		log.Fatal("debug argument is required")
	}

	return arguments{
		filePath:    *file,
		outPath:     *out,
		disassemble: *disassemble,
	}
}

func parseRunArgs() runArguments {
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatal("run requires a file argument")
	}

	return runArguments{
		filePath: os.Args[1],
	}
}
