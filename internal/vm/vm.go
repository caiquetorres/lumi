package vm

import "os"

type vm struct{}

func Execute(f *os.File) error {
	return nil
}

func New() *vm {
	return &vm{}
}
