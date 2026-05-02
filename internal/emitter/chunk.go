package emitter

import (
	"encoding/binary"
	"math"

	"github.com/caiquetorres/lumi/internal/constpool"
)

const defaultChunkSize = 512

type Chunk struct {
	ip uint32

	hasEntryPoint bool
	entryPoint    uint32

	code    []byte
	pool    *constpool.ConstantPool
	fnTable map[uint32]uint32
}

func newChunk() *Chunk {
	return &Chunk{
		code:    make([]byte, 0, defaultChunkSize),
		pool:    constpool.New(),
		fnTable: make(map[uint32]uint32),
	}
}

func (c *Chunk) emit(b byte) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, b)
	c.ip++
	return
}

func (c *Chunk) patch(offset uint32, b byte) {
	if offset < uint32(len(c.code)) {
		c.code[offset] = b
	}
}

func (c *Chunk) patchUint32(offset uint32, value uint32) {
	if offset+4 <= uint32(len(c.code)) {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], value)
		copy(c.code[offset:offset+4], buf[:])
	}
}

func (c *Chunk) emitBytes(bytes ...byte) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, bytes...)
	c.ip += uint32(len(bytes))
	return
}

func (c *Chunk) emitUint8(value uint8) (offset uint32) {
	offset = c.ip
	c.code = append(c.code, value)
	c.ip++
	return
}

func (c *Chunk) emitUint16(value uint16) (offset uint32) {
	offset = c.ip
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], value)

	c.code = append(c.code, buf[:]...)
	c.ip += 2
	return
}

func (c *Chunk) emitUint24(value uint32) (offset uint32) {
	offset = c.ip
	var buf [3]byte
	buf[0] = byte((value >> 16) & 0xFF)
	buf[1] = byte((value >> 8) & 0xFF)
	buf[2] = byte(value & 0xFF)

	c.code = append(c.code, buf[:]...)
	c.ip += 3
	return
}

func (c *Chunk) emitUint32(value uint32) (offset uint32) {
	offset = c.ip
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], value)

	c.code = append(c.code, buf[:]...)
	c.ip += 4
	return
}

func (c *Chunk) reserveUint32() (offset uint32) {
	offset = c.emitUint32(math.MaxUint32)
	return
}

func (c *Chunk) emitUint64(value uint64) (offset uint32) {
	offset = c.ip
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], value)

	c.code = append(c.code, buf[:]...)
	c.ip += 8
	return
}
