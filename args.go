package main

import (
	"flag"
	"log"
	"os"
)

type arguments struct {
	filePath string
	outPath  string
	debug    bool
}

type runArguments struct {
	execPath string
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

func parseRunArgs() runArguments {
	runFlags := flag.NewFlagSet("run", flag.ExitOnError)
	exec := runFlags.String("exec", "", "path to compiled bytecode file")

	if err := runFlags.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	if exec == nil || *exec == "" {
		log.Fatal("exec argument is required")
	}

	return runArguments{
		execPath: *exec,
	}
}
