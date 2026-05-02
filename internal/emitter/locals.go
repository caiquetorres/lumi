package emitter

type symbol struct {
	offset int
}

type locals struct {
	symbols   map[string]symbol
	parent    *locals
	nextIndex int
}

func newLocals(parent *locals) *locals {
	return &locals{
		symbols:   make(map[string]symbol),
		parent:    parent,
		nextIndex: 0,
	}
}

func (st *locals) define(name string) symbol {
	if idx, ok := st.symbols[name]; ok {
		return idx
	}

	idx := symbol{
		offset: st.nextIndex,
	}

	st.symbols[name] = idx
	st.nextIndex += 8

	return idx
}

func (st *locals) lookup(name string) (symbol, bool) {
	if sym, ok := st.symbols[name]; ok {
		return sym, true
	}

	if st.parent != nil {
		return st.parent.lookup(name)
	}

	return symbol{}, false
}
