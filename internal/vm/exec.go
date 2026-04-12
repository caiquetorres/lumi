package vm

import (
	"fmt"
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

	entryPoint, hasEntryPoint, err := getEntryPoint(src)
	if err != nil {
		return err
	}

	if !hasEntryPoint {
		return fmt.Errorf("no entry point found")
	}

	instructions, err := getInstructions(src)
	if err != nil {
		return err
	}

	machine := &vm{
		c:      c,
		src:    instructions,
		frames: []frame{{ptr: entryPoint}},
	}

	return machine.run()
}
