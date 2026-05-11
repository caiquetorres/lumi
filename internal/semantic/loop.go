package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Loop struct {
	Body *Block
}

func (t *TypeChecker) analyzeLoop(l *parser.Loop) *Loop {
	return &Loop{
		Body: t.analyzeBlock(l.Body),
	}
}
