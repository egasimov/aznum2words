package fixtures

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"testing"
)

type TestUseCaseHolder struct {
	UseCaseName string        `json:"useCaseName"`
	UseCases    []TestUseCase `json:"useCases"`
}

type TestUseCase struct {
	Name    string         `json:"name"`
	Samples []UseCaseDatum `json:"samples"`
}

type UseCaseDatum struct {
	Given    string `json:"given"`
	Expected string `json:"expected"`
}

func FixturePath(t *testing.T, fixtureName string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("problems recovering caller information")
	}

	// Get the absolute path of the current file
	absPath, err := filepath.Abs(filename)
	if err != nil {
		t.Fatal("Error getting absolute path:", err)
	}
	for i := 0; i < 2; i++ {
		absPath = filepath.Dir(absPath)
	}

	var result = filepath.Join(absPath, "fixtures", fixtureName)
	return result
}

func LoadTestCase(fileName string) (*TestUseCaseHolder, error) {
	var testUseCaseHolder TestUseCaseHolder
	rawBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("Error occurred while reading config")
		return nil, err
	}

	err = json.Unmarshal(rawBytes, &testUseCaseHolder)

	if err != nil {
		return nil, err
	}

	return &testUseCaseHolder, nil
}
