package generator

import "errors"

func init() {
	addGenerator("go", newGoGenerator)
}

func newGoGenerator(opts map[string]string) (Generator, error) {
	return &GoGenerator{}, nil
}

type GoGenerator struct {
}

func (g *GoGenerator) Generate(sw *Swagger) error {
	return errors.New("Not Implemented")
}
