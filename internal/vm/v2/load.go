package vm

import (
	"errors"
	"io"

	"github.com/caiquetorres/lumi/internal/emitter"
)

func (m *vm) load(entryPoint uint32) error {
	m.frames.push(entryPoint, 0, m.src)

	for {
		opcode, err := m.frames.current().readUint8(m.src)
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}

		switch opcode {
		case emitter.FnDecl:
			_, _ = m.frames.current().readUint32(m.src) // name index
			_, _ = m.frames.current().readUint32(m.src) // entry point

		case emitter.PushInt:
			_, _ = m.frames.current().readUint32(m.src)

		case emitter.PushTrue, emitter.PushFalse, emitter.Add,
			emitter.Return, emitter.Pop:
		}
	}
}
