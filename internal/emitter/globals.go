package emitter

type globals struct {
	symbols map[string]uint32
}

func newGlobals() *globals {
	return &globals{
		symbols: make(map[string]uint32),
	}
}

func (g *globals) len() int {
	return len(g.symbols)
}

func (g *globals) define(name string) uint32 {
	if idx, ok := g.symbols[name]; ok {
		return idx
	}

	idx := uint32(len(g.symbols))
	g.symbols[name] = idx

	return idx
}

func (g *globals) lookup(name string) (uint32, bool) {
	idx, ok := g.symbols[name]
	return idx, ok
}
