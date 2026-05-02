package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *Emitter) BeforeBlockStmt(block *parser.Block) {}

func (e *Emitter) AfterBlockStmt(block *parser.Block) {}
