package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type FunDecl struct {
	Identifier token.Token
	Params     []Param
	Return     *Type
	Body       []Stmt
}

func (a *Analyzer) analyzeFunDecl(fd *parser.FunDecl) *FunDecl {
	params := make([]Param, len(fd.Params))
	for i, p := range fd.Params {
		params[i] = a.analyzeParam(p)
	}

	body := make([]Stmt, len(fd.Body))
	for i, s := range fd.Body {
		body[i] = a.analyzeStmt(s)
	}

	return &FunDecl{
		Identifier: fd.Identifier,
		Params:     params,
		Body:       body,
		Return:     a.analyzeType(fd.Return),
	}
}

type Param struct {
	Name token.Token
	Type *Type
}

func (a *Analyzer) analyzeParam(p parser.Param) Param {
	return Param{
		Name: p.Name,
		Type: a.analyzeType(p.Type),
	}
}
