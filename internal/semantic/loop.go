package semantic

import "github.com/caiquetorres/lumi/internal/parser"

type Loop struct {
	Body *Block
}

func (a *Analyzer) analyzeLoop(l *parser.Loop) *Loop {
	return &Loop{
		Body: a.analyzeBlock(l.Body),
	}
}
