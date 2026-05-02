package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeBlockStmt(block *parser.Block) {}

func (e *emitter) AfterBlockStmt(block *parser.Block) {}
