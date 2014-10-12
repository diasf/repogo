package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func ParseSwagger(file string) (*Swagger, error) {
	var f *os.File
	var err error
	if f, err = os.Open(file); err != nil {
		return nil, FileNotFoundError{FileError{file, ParentError{err}}}
	}
	swagger := &Swagger{}
	log.Printf("Going to parse swagger file %s", file)
	if err = json.NewDecoder(f).Decode(swagger); err != nil {
		return nil, errors.New("Error unmarshaling json: " + err.Error())
	}
	return swagger, nil
}

type FileNotFoundError struct {
	FileError
}

func (f FileNotFoundError) Error() string {
	return fmt.Sprintf("Error reading file", f.FileError.Error())
}

type UnmarshallError struct {
	FileError
}

func (u UnmarshallError) Error() string {
	return fmt.Sprintf("Error unmarshaling json", u.FileError.Error())
}

type FileError struct {
	File string
	ParentError
}

func (f FileError) Error() string {
	return fmt.Sprintf("[%s: %s]", f.File, f.ParentError.Error())
}
