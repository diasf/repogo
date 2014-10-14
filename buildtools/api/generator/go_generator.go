package generator

func init() {
	addGenerator("go", newGoGenerator)
}

func newGoGenerator(opts map[string]string) (Generator, error) {
	return &GoGenerator{}, nil
}

type GoGenerator struct {
}

func (g *GoGenerator) Generate(sw *Swagger) error {
	return nil //errors.New("Not Implemented")
}
