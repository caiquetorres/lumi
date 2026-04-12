package vm

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type constantPool struct {
	constants  []any
	indexByKey map[string]int
}

func parseConstantPool(data []byte) (*constantPool, error) {
	pool := &constantPool{
		indexByKey: map[string]int{},
	}

	if len(data) == 0 {
		return pool, nil
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		idx, value, key, err := parseConstantLine(line)
		if err != nil {
			return nil, err
		}

		for len(pool.constants) <= idx {
			pool.constants = append(pool.constants, nil)
		}

		pool.constants[idx] = value
		pool.indexByKey[key] = idx
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pool, nil
}

func parseConstantLine(line string) (int, any, string, error) {
	if !strings.HasPrefix(line, "#") {
		return 0, nil, "", fmt.Errorf("invalid constant line: %q", line)
	}

	parts := strings.SplitN(line[1:], ": ", 2)
	if len(parts) != 2 {
		return 0, nil, "", fmt.Errorf("invalid constant line: %q", line)
	}

	idx, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, nil, "", fmt.Errorf("invalid constant index %q: %w", parts[0], err)
	}

	typeAndValue := strings.SplitN(parts[1], ":", 2)
	if len(typeAndValue) != 2 {
		return 0, nil, "", fmt.Errorf("invalid constant payload: %q", parts[1])
	}

	typ := typeAndValue[0]
	rawValue := typeAndValue[1]

	switch typ {
	case "int":
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return 0, nil, "", fmt.Errorf("invalid int constant %q: %w", rawValue, err)
		}

		return idx, value, strconv.Itoa(value), nil

	case "string":
		value, err := strconv.Unquote(rawValue)
		if err != nil {
			return 0, nil, "", fmt.Errorf("invalid string constant %q: %w", rawValue, err)
		}

		return idx, value, value, nil

	default:
		return 0, nil, "", fmt.Errorf("unsupported constant type %q", typ)
	}
}

func (c *constantPool) getConstant(index int) (any, bool) {
	if index < len(c.constants) {
		return c.constants[index], true
	}
	return nil, false
}

func (c *constantPool) getIndex(key string) (int, bool) {
	if c.indexByKey == nil {
		return 0, false
	}
	idx, ok := c.indexByKey[key]
	return idx, ok
}
