package emitter

import (
	"bufio"
	"encoding/binary"
	"io"
)

const lumiMagic = "LUMI"

type builder struct {
	w *bufio.Writer
}

func newBuilder(w io.Writer) *builder {
	return &builder{w: bufio.NewWriter(w)}
}

func (b *builder) build(ch *Chunk) error {
	if err := b.writeMagic(); err != nil {
		return err
	}

	if err := b.writeConstantPool(ch.pool.Serialize()); err != nil {
		return err
	}

	if ch.hasEntryPoint {
		if err := b.w.WriteByte(1); err != nil {
			return err
		}
		if err := b.writeUint32(ch.entryPoint); err != nil {
			return err
		}
	} else {
		if err := b.w.WriteByte(10); err != nil {
			return err
		}
	}

	if _, err := b.w.Write(ch.code); err != nil {
		return err
	}

	return b.w.Flush()
}

func (b *builder) writeMagic() error {
	_, err := b.w.WriteString(lumiMagic)
	return err
}

func (b *builder) writeConstantPool(constantPool []byte) error {
	if err := b.writeUint32(uint32(len(constantPool))); err != nil {
		return err
	}

	_, err := b.w.Write(constantPool)
	return err
}

func (b *builder) writeUint32(value uint32) error {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], value)

	_, err := b.w.Write(buf[:])
	return err
}
