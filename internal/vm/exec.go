package vm

import (
	"io"

	"github.com/caiquetorres/lumi/internal/constpool"
)

func Execute(src io.ReadSeeker) error {
	if !isLumiFile(src) {
		return nil
	}

	constants, err := getConstants(src)
	if err != nil {
		return err
	}

	pool, err := constpool.ParseConstantPool(constants)
	if err != nil {
		return err
	}

	instructions, err := getInstructions(src)
	if err != nil {
		return err
	}

	globals := newGlobalSymbolTable()

	machine := &vm{
		pool:        pool,
		src:         instructions,
		globals:     globals,
		symbolTable: globals,
		stack:       make([]any, 0, 16),
	}

	if err := machine.load(); err != nil {
		return err
	}

	return machine.run()
}
