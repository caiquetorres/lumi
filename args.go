package main

import (
	"flag"
	"log"
)

type arguments struct {
	filePath string
	outPath  string
	debug    bool
}

func parseArgs() arguments {
	file := flag.String("file", "", "path to source file")
	out := flag.String("out", "", "path to output file")
	debug := flag.Bool("debug", false, "enable debug mode")

	flag.Parse()

	if file == nil || *file == "" {
		log.Fatal("file argument is required")
	}

	if out == nil || *out == "" {
		log.Fatal("out argument is required")
	}

	if debug == nil {
		log.Fatal("debug argument is required")
	}

	return arguments{
		filePath: *file,
		outPath:  *out,
		debug:    *debug,
	}
}
