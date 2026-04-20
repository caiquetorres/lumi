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

	machine := &vm{
		c:       c,
		src:     instructions,
		globals: newSymbolTable(nil),
	}

	if err := machine.load(); err != nil {
		return err
	}

	return machine.run()
}
