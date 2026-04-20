package parser

// TODO: We don't need to make the circular buffer generic and with variable size, we can just make it a fixed-size buffer of tokens. This will simplify the implementation and reduce the memory usage.

type circularBuffer struct {
	size uint32
	head uint32
	len  uint32
	buf  []any
}

func newCircularBuffer(size uint32) *circularBuffer {
	return &circularBuffer{
		size: size,
		buf:  make([]any, size),
	}
}

func (c *circularBuffer) pushBack(v any) {
	tail := (c.head + c.len) % c.size
	c.buf[tail] = v
	c.len++
}

func (c *circularBuffer) popFront() any {
	v := c.buf[c.head]
	c.head = (c.head + 1) % c.size
	c.len--
	return v
}

func (c *circularBuffer) at(n uint32) any {
	return c.buf[(c.head+n)%c.size]
}

func (c *circularBuffer) isFull() bool {
	return c.len == c.size
}

func (c *circularBuffer) isEmpty() bool {
	return c.len == 0
}
