package generator

import "testing"

func TestExecuteSuccess(t *testing.T) {
	if err := Execute(PETSTORE_EXPANDED_FILE, "go", nil); err != nil {
		t.Errorf("Execution with errors", err.Error())
	}
}

func TestExecuteInvalidGenerator(t *testing.T) {
	if err := Execute(PETSTORE_EXPANDED_FILE, "INVALID GENERATOR", nil); err == nil {
		t.Errorf("Execution should have failed with invalid generator")
	} else if _, ok := err.(GeneratorNotFoundError); !ok {
		t.Errorf("Expected GeneratorNotFoundError")
	}
}
