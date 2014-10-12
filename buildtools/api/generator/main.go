package generator

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func ParseSwagger(file string) (*Swagger, error) {
	var f *os.File
	var err error
	if f, err = os.Open(file); err != nil {
		return nil, errors.New("Error reading file: " + err.Error())
	}
	swagger := &Swagger{}
	log.Printf("Going to parse swagger file %s", file)
	if err = json.NewDecoder(f).Decode(swagger); err != nil {
		return nil, errors.New("Error unmarshaling json: " + err.Error())
	}
	return swagger, nil
}
