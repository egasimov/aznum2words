//go:build integration
// +build integration

package it

import (
	"context"
	"encoding/json"
	"fmt"
	aznum2wordsclient "github.com/egasimov/aznum2words/cmd/aznum2words-webapp/it/aznum2wordsclient"
	"github.com/egasimov/aznum2words/fixtures"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestConverterAPI_Success(t *testing.T) {
	testDataFileNames := []string{
		"positive-integers.json",
		"negative-integers.json",
		"positive-floating-point-numbers.json",
		"negative-floating-point-numbers.json",
	}

	appHostVal := os.Getenv("APP_HOST")
	appPortVal := os.Getenv("APP_PORT")
	if appPortVal == "" {
		appPortVal = "8080"
	}

	metricPortVal := os.Getenv("METRIC_PORT")
	if metricPortVal == "" {
		metricPortVal = "9090"
	}

	apiClient := aznum2wordsclient.AzNum2WordsClient{
		Server:         fmt.Sprintf("http://%s:%s", appHostVal, appPortVal),
		Client:         http.DefaultClient,
		RequestEditors: nil,
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
					for idx, sampleItem := range useCase.Samples {
						requestBody := aznum2wordsclient.ConvertNumberToWordsRequest{
							Number: sampleItem.Given,
						}

						resp, err := apiClient.ConvertNumberToWord(
							context.Background(),
							&aznum2wordsclient.ConvertNumberToWordParams{},
							requestBody,
						)
						if err != nil {
							t.Fatal(err)
						}

						resBody, err := ioutil.ReadAll(resp.Body)

						var response = aznum2wordsclient.ConvertWordsToNumber{}

						if err := json.Unmarshal(resBody, &response); err != nil {
							t.Fatal(err)
						}

						actual := response.Words
						if !reflect.DeepEqual(actual, sampleItem.Expected) {
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
