package emitter

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/caiquetorres/lumi/internal/constpool"
)

func (c *Chunk) Serialize(w io.Writer) error {
	b := bufio.NewWriter(w)

	if err := writeConstantPool(w, c.pool); err != nil {
		return err
	}

	if err := writeFunctionTable(w, c.fnTable); err != nil {
		return err
	}

	if c.hasEntryPoint {
		if err := b.WriteByte(1); err != nil {
			return err
		}
		if err := writeUint32(b, c.entryPoint); err != nil {
			return err
		}
	} else {
		if err := b.WriteByte(0); err != nil {
			return err
		}
	}

	if _, err := b.Write(c.code); err != nil {
		return err
	}

	return nil
}

func writeFunctionTable(w io.Writer, fnTable map[uint32]uint32) error {
	if err := writeUint32(w, uint32(len(fnTable))); err != nil {
		return err
	}

	for id, offset := range fnTable {
		if err := writeUint32(w, id); err != nil {
			return err
		}
		if err := writeUint32(w, offset); err != nil {
			return err
		}
	}

	return nil
}

func writeConstantPool(w io.Writer, constantPool *constpool.ConstantPool) error {
	buf := constantPool.Serialize()

	if err := writeUint32(w, uint32(len(buf))); err != nil {
		return err
	}

	_, err := w.Write(buf)
	return err
}

func writeUint32(w io.Writer, value uint32) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], value)

	_, err := w.Write(buf[:])
	return err
}
