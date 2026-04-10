package main

import (
	"flag"
	"log"
)

type arguments struct {
	filePath string
}

func parseArgs() arguments {
	file := flag.String("file", "", "path to source file")
	flag.Parse()

	if file == nil || *file == "" {
		log.Panic("file argument is required")
	}

	return arguments{
		filePath: *file,
	}
}
