package vm

import (
	"io"
)

func Execute(src io.ReadSeeker) error {
	if !isLumiFile(src) {
		return nil
	}

	constants, err := getConstants(src)
	if err != nil {
		return err
	}

	c, err := parseConstantPool(constants)
	if err != nil {
		return err
	}

	instructions, err := getInstructions(src)
	if err != nil {
		return err
	}

	globals := newSymbolTable(nil)

	machine := &vm{
		c:           c,
		src:         instructions,
		globals:     globals,
		symbolTable: globals,
	}

	if err := machine.load(); err != nil {
		return err
	}

	return machine.run()
}
