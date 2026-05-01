package vm

type symbolTable struct {
	table map[string]any

	parent *symbolTable
}

func newGlobalSymbolTable() *symbolTable {
	return newSymbolTable(nil)
}

func newSymbolTable(parent *symbolTable) *symbolTable {
	return &symbolTable{
		table:  make(map[string]any),
		parent: parent,
	}
}

func (s *symbolTable) define(name string, value any) {
	s.table[name] = value
}

func (s *symbolTable) set(name string, value any) bool {
	if _, exists := s.table[name]; exists {
		s.table[name] = value
		return true
	}

	if s.parent != nil {
		return s.parent.set(name, value)
	}

	return false
}

func (s *symbolTable) lookup(name string) (any, bool) {
	val, exists := s.table[name]
	if !exists && s.parent != nil {
		return s.parent.lookup(name)
	}

	return val, exists
}
