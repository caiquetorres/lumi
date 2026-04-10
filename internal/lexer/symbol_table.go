package lexer

import "github.com/caiquetorres/lumi/internal/token"

type symbolTable struct {
	symbols []string
	index   map[string]token.SymbolID
}

func newSymbolTable() *symbolTable {
	return &symbolTable{
		symbols: []string{},
		index:   make(map[string]token.SymbolID),
	}
}

func (st *symbolTable) intern(s string) token.SymbolID {
	if id, ok := st.index[s]; ok {
		return id
	}

	id := token.SymbolID(len(st.symbols))
	st.symbols = append(st.symbols, s)
	st.index[s] = id

	return id
}

func (st *symbolTable) lookup(id token.SymbolID) string {
	return st.symbols[id]
}
