package vm

type symbolTable struct {
	table map[string]any

	parent *symbolTable
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

func (s *symbolTable) lookup(name string) (any, bool) {
	val, exists := s.table[name]
	if !exists && s.parent != nil {
		return s.parent.lookup(name)
	}

	return val, exists
}
