package semantic

type SymbolTable struct {
	types  map[string]Kind
	parent *SymbolTable
}

func NewRootSymbolTable() *SymbolTable {
	return NewSymbolTable(nil)
}

func NewSymbolTable(parent *SymbolTable) *SymbolTable {
	return &SymbolTable{
		types:  make(map[string]Kind),
		parent: parent,
	}
}

func (s *SymbolTable) Define(name string, kind Kind) {
	s.types[name] = kind
}

func (s *SymbolTable) Lookup(name string) (Kind, bool) {
	if kind, ok := s.types[name]; ok {
		return kind, true
	}

	if s.parent != nil {
		return s.parent.Lookup(name)
	}

	return nil, false
}
