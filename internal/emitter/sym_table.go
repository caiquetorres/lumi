package emitter

type symbol struct {
	offset, size int

	ty     string
	isHeap bool
}

type symbolTable struct {
	symbols   map[string]symbol
	parent    *symbolTable
	nextIndex int
}

func newSymbolTable(parent *symbolTable, nextIndex int) *symbolTable {
	return &symbolTable{
		symbols:   make(map[string]symbol),
		parent:    parent,
		nextIndex: nextIndex,
	}
}

func (st *symbolTable) define(name, ty string, size int, isHeap bool) symbol {
	if idx, ok := st.symbols[name]; ok {
		return idx
	}

	idx := symbol{
		offset: st.nextIndex,
		size:   size,
		ty:     ty,
		isHeap: isHeap,
	}

	st.symbols[name] = idx
	st.nextIndex += size

	return idx
}

func (st *symbolTable) lookup(name string) (symbol, bool) {
	if sym, ok := st.symbols[name]; ok {
		return sym, true
	}

	if st.parent != nil {
		return st.parent.lookup(name)
	}

	return symbol{}, false
}
