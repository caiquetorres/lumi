package parser

type parseError struct {
	errs    []error
	isFatal bool
}

func (p *parseError) Error() string {
	panic("unimplemented")
}
