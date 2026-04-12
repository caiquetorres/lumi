package vm

import (
	"encoding/binary"
	"io"
)

func isLumiFile(fp io.ReadSeeker) bool {
	magic := make([]byte, 4)

	n, err := fp.Read(magic)
	if err != nil || n != 4 {
		return false
	}

	return string(magic) == "LUMI"
}

func getConstantsSize(fp io.ReadSeeker) (int, error) {
	sizeBuf := make([]byte, 4)

	n, err := fp.Read(sizeBuf)
	if err != nil || n != 4 {
		return 0, err
	}

	return int(binary.BigEndian.Uint32(sizeBuf)), nil
}

func getConstants(fp io.ReadSeeker) ([]byte, error) {
	size, err := getConstantsSize(fp)
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

func getEntryPoint(fp io.ReadSeeker) (int, bool, error) {
	entryPointFlag := make([]byte, 1)

	n, err := fp.Read(entryPointFlag)
	if err != nil || n != 1 {
		return 0, false, err
	}

	if entryPointFlag[0] == 0 {
		return 0, false, nil
	}

	entryPointBuf := make([]byte, 4)

	n, err = fp.Read(entryPointBuf)
	if err != nil || n != 4 {
		return 0, false, err
	}

	return int(binary.BigEndian.Uint32(entryPointBuf)), true, nil
}

func getInstructions(fp io.ReadSeeker) ([]byte, error) {
	instructions, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	return instructions, nil
}
