package main

import (
	"io"
	"log"
	"os"

	"github.com/caiquetorres/lumi/internal/emitter"
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/semantic"
	"github.com/caiquetorres/lumi/internal/vm"
)

func main() {
	args := parseArgs()
	log.Printf("running with file: %s", args.filePath)

	f, err := os.Open(args.filePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := runPipeline(f); err != nil {
		log.Fatal(err)
	}
}

func runPipeline(r io.Reader) error {
	ast, parseErr := parser.Parse(r)
	if parseErr != nil {
		return parseErr
	}

	semanticErr := semantic.Analyze(ast)
	if semanticErr != nil {
		return semanticErr
	}

	emitErr := emitter.Emit(ast, os.Stdout)
	if emitErr != nil {
		return emitErr
	}

	executionErr := vm.Execute()
	if executionErr != nil {
		return executionErr
	}

	return nil
}
