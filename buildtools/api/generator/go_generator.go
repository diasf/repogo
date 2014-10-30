package generator

import "errors"

func init() {
	addGenerator("go", newGoGenerator)
}

func newGoGenerator(opts map[string]string) (Generator, error) {
	return &GoGenerator{}, nil
}

type GoGenerator struct {
	Dest string
}

func (g *GoGenerator) Generate(sw *Swagger) error {
	if sw == nil {
		return errors.New("Invalid parameter")
	}

	return nil
}
