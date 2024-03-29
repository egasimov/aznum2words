//go:build integration
// +build integration

package main

import (
	"flag"
	"fmt"
	"github.com/egasimov/aznum2words/fixtures"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"reflect"
)

var failFastFlag = flag.Bool("failfast", false, "stops in case of first test failure")

var binaryName = "aznum2words-cli"

func TestCliArgs(t *testing.T) {

	testDataFileNames := []string{
		"positive-integers.json",
		"negative-integers.json",
		"positive-floating-point-numbers.json",
		"negative-floating-point-numbers.json",
	}

	for _, fileName := range testDataFileNames {
		// load file containing samples for a  use case
		testUseCaseHolder, err := fixtures.LoadTestCase(fixtures.FixturePath(t, fileName))

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
						cmd := exec.Command(path.Join(dir, binaryName), "--", sampleItem.Given)
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
