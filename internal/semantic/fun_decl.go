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

func funDecl(idenfier token.Token, params []Param, returnType *Type, body []Stmt) *FunDecl {
	return &FunDecl{
		Identifier: idenfier,
		Params:     params,
		Return:     returnType,
		Body:       body,
	}
}

func (a *TypeChecker) analyzeFunDecl(fd *parser.FunDecl) *FunDecl {
	params := make([]Param, len(fd.Params))
	for i, p := range fd.Params {
		params[i] = a.analyzeParam(p)
	}

	re := a.analyzeType(fd.Return)

	body := make([]Stmt, len(fd.Body))
	for i, s := range fd.Body {
		body[i] = a.analyzeStmt(s)
	}

	sre := a.parseType(fd.Return)
	sparams := make([]Kind, len(params))
	for i, p := range fd.Params {
		sparams[i] = a.parseType(p.Type)
	}

	name := a.lex.Lexeme(fd.Identifier)
	a.symTable.Define(name, Function{
		Params: sparams,
		Return: sre,
	})

	return funDecl(fd.Identifier, params, re, body)
}

type Param struct {
	Name token.Token
	Type *Type
}

func (a *TypeChecker) analyzeParam(p parser.Param) Param {
	return Param{
		Name: p.Name,
		Type: a.analyzeType(p.Type),
	}
}
