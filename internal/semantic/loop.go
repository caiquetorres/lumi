package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Loop struct {
	Body *Block
}

func loop(l *parser.Loop) *Loop {
	return &Loop{
		Body: block(l.Body),
	}
}
