package vm

import (
	"encoding/binary"
	"io"
)

const lumiMagic = "LUMI"

func isLumiFile(fp io.ReadSeeker) bool {
	var magic [4]byte

	n, err := fp.Read(magic[:])
	if err != nil || n != 4 {
		return false
	}

	return string(magic[:]) == lumiMagic
}

func getConstantPoolLen(fp io.ReadSeeker) (int, error) {
	var sizeBuf [4]byte

	n, err := fp.Read(sizeBuf[:])
	if err != nil || n != 4 {
		return 0, err
	}

	return int(binary.BigEndian.Uint32(sizeBuf[:])), nil
}

func getConstants(fp io.ReadSeeker) ([]byte, error) {
	size, err := getConstantPoolLen(fp)
	if err != nil {
		return nil, err
	}

	constants := make([]byte, size)

	n, err := fp.Read(constants)
	if err != nil || n != size {
		return nil, err
	}

	return constants, nil
}

func getInstructions(fp io.ReadSeeker) ([]byte, error) {
	instructions, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}
