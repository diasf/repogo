package generator

import "fmt"

func Execute(schemaFile, generatorId string, options map[string]string) error {
	var err error
	var sw *Swagger
	var generator Generator
	if sw, err = ParseSwagger(schemaFile); err != nil {
		return err
	}

	genFunc := generators[generatorId]
	if genFunc == nil {
		return GeneratorNotFoundError{generatorId}
	}

	if generator, err = genFunc(options); err != nil {
		return GeneratorInitialisationError{generatorId, ParentError{err}}
	}

	return generator.Generate(sw)
}

type generatorInitialiser func(opts map[string]string) (Generator, error)

var generators = make(map[string]generatorInitialiser)

func addGenerator(id string, gen generatorInitialiser) {
	generators[id] = gen
}

type Generator interface {
	Generate(sw *Swagger) error
}

type GeneratorNotFoundError struct {
	Id string
}

func (g GeneratorNotFoundError) Error() string {
	return fmt.Sprintf("Could not find generator '%s'", g.Id)
}

type GeneratorInitialisationError struct {
	Id string
	ParentError
}

func (g GeneratorInitialisationError) Error() string {
	return fmt.Sprintf("Could not initialize generator '%s': %s", g.Id, g.ParentError.Error())
}

type ParentError struct {
	Parent error
}

func (p ParentError) Error() string {
	var parentError string
	if p.Parent != nil {
		parentError = p.Parent.Error()
	}
	return parentError
}
