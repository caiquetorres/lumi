package semantic

import (
	"github.com/caiquetorres/lumi/internal/parser"
	"github.com/caiquetorres/lumi/internal/token"
)

type Fun struct {
	Identifier token.Token
	Params     []Param
	Return     *Type
	Body       []Stmt
}

func funDecl(
	idenfier token.Token, params []Param,
	returnType *Type, body []Stmt,
) *Fun {
	return &Fun{
		Identifier: idenfier,
		Params:     params,
		Return:     returnType,
		Body:       body,
	}
}

func (t *TypeChecker) analyzeFunDecl(fd *parser.FunDecl) *Fun {
	params := make([]Param, len(fd.Params))
	for i, p := range fd.Params {
		params[i] = t.analyzeParam(p)
	}

	re := t.analyzeType(fd.Return)

	body := make([]Stmt, len(fd.Body))
	for i, s := range fd.Body {
		body[i] = t.analyzeStmt(s)
	}

	sre := t.parseType(fd.Return)
	sparams := make([]Kind, len(params))
	for i, p := range fd.Params {
		sparams[i] = t.parseType(p.Type)
	}

	name := t.lex.Lexeme(fd.Identifier)
	t.symTable.Define(name, Function{
		Params: sparams,
		Return: sre,
	})

	return funDecl(fd.Identifier, params, re, body)
}

type Param struct {
	Name token.Token
	Type *Type
}

func (t *TypeChecker) analyzeParam(p parser.Param) Param {
	return Param{
		Name: p.Name,
		Type: t.analyzeType(p.Type),
	}
}
