package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"reflect"
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

var failFastFlag = flag.Bool("failfast", false, "stops in case of first test failure")

var binaryName = "aznum2words-cli"

func fixturePath(t *testing.T, fixture string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("problems recovering caller information")
	}

	return filepath.Join(filepath.Dir(filename), fixture)
}

func TestCliArgs(t *testing.T) {

	testDataFileNames := []string{
		"positive-integers.json",
		"negative-integers.json",
		"positive-floating-point-numbers.json",
		"negative-floating-point-numbers.json",
	}

	for _, fileName := range testDataFileNames {
		// load file containing samples for a  use case
		testUseCaseHolder, err := loadTestCase(fixturePath(t, fileName))

		if err != nil {
			t.Fatal(err)
		}
		for _, useCase := range testUseCaseHolder.UseCases {
			t.Run(testUseCaseHolder.UseCaseName+"$"+useCase.Name,
				func(t *testing.T) {
					dir, err := os.Getwd()
					if err != nil {
						t.Fatal(err)
					}

					for idx, sampleItem := range useCase.Samples {
						args := []string{"\"" + sampleItem.Given + "\""}
						cmd := exec.Command(path.Join(dir, binaryName), strings.Join(args, ""))
						outputRawByte, err := cmd.CombinedOutput()

						if err != nil {
							fmt.Println("Executed Command: " + cmd.String())
							fmt.Println("Console Output: " + string(outputRawByte))
							t.Fatal(err)
						}

						actual := string(outputRawByte)

						//TODO when reading from output from console, add space to end of real output
						actual = strings.TrimSpace(actual)
						if !reflect.DeepEqual(actual, sampleItem.Expected) {
							if *failFastFlag {
								t.Fatalf("For %s samples[%d]"+
									"\n Given: %s, len: %d "+
									"\n Expected: %s, len: %d"+
									"\n Got: %s, len: %d",
									t.Name(), idx,
									sampleItem.Given, len(sampleItem.Given),
									sampleItem.Expected, len(sampleItem.Expected),
									actual, len(actual))
							}
							t.Fatalf("For %s samples[%d]"+
								"\n Given: %s, len: %d "+
								"\n Expected: %s, len: %d"+
								"\n Got: %s, len: %d",
								t.Name(), idx,
								sampleItem.Given, len(sampleItem.Given),
								sampleItem.Expected, len(sampleItem.Expected),
								actual, len(actual))
						}
					}
				})
		}

	}
}

func testUseCase(name string, useCase TestUseCase) {

}

func TestMain(m *testing.M) {
	err := os.Chdir("../..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}
	make := exec.Command("make")
	err = make.Run()
	if err != nil {
		fmt.Printf("could not make binary for %s: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func loadTestCase(fileName string) (*TestUseCaseHolder, error) {
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
