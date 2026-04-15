package emitter

import "github.com/caiquetorres/lumi/internal/parser"

func (e *emitter) BeforeVarDecl(vd *parser.VarDecl) error {
	return nil
}

func (e *emitter) AfterVarDecl(vd *parser.VarDecl) error {
	if err := e.emit(VarDecl); err != nil {
		return err
	}

	name := e.l.Lexeme(vd.Identifier)
	if err := e.writeStringConstant(name); err != nil {
		return err
	}

	return e.flush()
}
