package parser

import (
	"errors"
)

var (
	ErrUnexpectedToken = errors.New("unexpected token")
	ErrUnexpectedEOF   = errors.New("unexpected end of file")
)
