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

func (b *builder) build(constantPool []byte, instructions []byte) error {
	if err := b.writeMagic(); err != nil {
		return err
	}

	if err := b.writeConstantPool(constantPool); err != nil {
		return err
	}

	if _, err := b.w.Write(instructions); err != nil {
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
