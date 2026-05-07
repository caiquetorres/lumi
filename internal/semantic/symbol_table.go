package semantic

type SymbolTable struct {
	types  map[string]Kind
	parent *SymbolTable
}
