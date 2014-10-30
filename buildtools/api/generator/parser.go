package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ParseSwagger(file string) (*Swagger, error) {
	var f *os.File
	var err error
	if f, err = os.Open(file); err != nil {
		return nil, FileNotFoundError{file, err}
	}
	swagger := &Swagger{}
	log.Printf("Going to parse swagger file %s", file)
	if err = json.NewDecoder(f).Decode(swagger); err != nil {
		return nil, UnmarshallError{file, err}
	}
	return swagger, nil
}

type FileNotFoundError struct {
	File string
	Err  error
}

func (f FileNotFoundError) Error() string {
	return fmt.Sprintf("Error reading file [%s: %s]", f.File, f.Err.Error())
}

type UnmarshallError struct {
	File string
	Err  error
}

func (u UnmarshallError) Error() string {
	return fmt.Sprintf("Error unmarshaling json [%s: %s]", u.File, u.Err.Error())
}
