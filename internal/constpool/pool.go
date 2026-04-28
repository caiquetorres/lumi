package constpool

import "fmt"

type ConstantPool struct {
	constants  []any
	indexByKey map[string]uint32
}

func New() *ConstantPool {
	return &ConstantPool{
		constants:  make([]any, 0),
		indexByKey: make(map[string]uint32),
	}
}

func (c *ConstantPool) InternConstant(value any) uint32 {
	if c.indexByKey == nil {
		c.indexByKey = make(map[string]uint32)
	}

	key := constantKey(value)
	if idx, ok := c.indexByKey[key]; ok {
		return idx
	}

	c.constants = append(c.constants, value)
	idx := uint32(len(c.constants) - 1)
	c.indexByKey[key] = idx

	return idx
}

func (c *ConstantPool) GetConstant(index uint32) (any, bool) {
	if index < uint32(len(c.constants)) {
		return c.constants[index], true
	}

	return nil, false
}

func (c *ConstantPool) GetConstantAsString(index uint32) (string, error) {
	value, exists := c.GetConstant(index)
	if !exists {
		return "", fmt.Errorf("constant with index %d not found", index)
	}

	strValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("expected string constant at index %d, got %T", index, value)
	}

	return strValue, nil
}

func constantKey(value any) string {
	return fmt.Sprintf("%T:%#v", value, value)
}
