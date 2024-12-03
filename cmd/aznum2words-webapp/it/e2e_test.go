//go:build integration
// +build integration

package it

import (
	"context"
	"encoding/json"
	"fmt"
	aznum2wordsclient "github.com/egasimov/aznum2words/cmd/aznum2words-webapp/it/aznum2wordsclient"
	"github.com/egasimov/aznum2words/fixtures"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type convertNumber2WordsEndToEndTests struct {
	*IntegrationTestFixture
}

type IntegrationTestFixture struct {
	*testing.T
	Ctx    context.Context
	WebApp *testcontainers.Container
}

func NewIntegrationTestFixture(t *testing.T) *IntegrationTestFixture {
	ctx := context.Background()

	// Build docker image and run the container
	app, err := setupWithDockerFile(t, ctx, "../../../", "webapp.Dockerfile")
	if err != nil {
		t.Fatal(err)
	}

	integrationTestFixture := &IntegrationTestFixture{T: t}
	integrationTestFixture.WebApp = &app
	integrationTestFixture.Ctx = ctx

	return integrationTestFixture
}

func TestRunner(t *testing.T) {
	var endToEndTestFixture = NewIntegrationTestFixture(t)

	//https://pkg.go.dev/testing@master#hdr-Subtests_and_Sub_benchmarks
	t.Run("A=convert-number-to-words-api-testing", func(t *testing.T) {

		testFixture := &convertNumber2WordsEndToEndTests{endToEndTestFixture}
		testFixture.Test_Should_Return_Success_When_Convert_Number2Words_Called()
	})
}

func (c *convertNumber2WordsEndToEndTests) Test_Should_Return_Success_When_Convert_Number2Words_Called() {
	appHostVal, appPortVal, _, err := getHostAndPorts(c.Ctx, *c.WebApp)
	if err != nil {
		c.T.Fatal(err)
	}

	apiClient := aznum2wordsclient.AzNum2WordsClient{
		Server:         fmt.Sprintf("http://%s:%s", appHostVal, appPortVal),
		Client:         http.DefaultClient,
		RequestEditors: nil,
	}

	testDataFileNames := []string{
		"positive-integers.json",
		"negative-integers.json",
		"positive-floating-point-numbers.json",
		"negative-floating-point-numbers.json",
	}

	for _, fileName := range testDataFileNames {
		// load file containing samples for a  use case
		testUseCaseHolder, err := fixtures.LoadTestCase(fixtures.FixturePath(c.T, fileName))

		if err != nil {
			c.T.Fatal(err)
		}
		for _, useCase := range testUseCaseHolder.UseCases {
			c.T.Run(testUseCaseHolder.UseCaseName+"$"+useCase.Name,
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

func setupWithDockerFile(t *testing.T, ctx context.Context, dockerfileCtx string, dockerfileName string) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    dockerfileCtx,
			Dockerfile: dockerfileName,
			Repo:       "local/github.com/egasimov/aznum2words",
			Tag:        "webapp-e2e",
			KeepImage:  false,
		},
		ExposedPorts: []string{"8080/tcp", "9090/tcp"}, //container port to expose
		Env:          map[string]string{                //Values for container environmental variables
		},
		//Checking for APP started listening on this port
		WaitingFor: wait.ForListeningPort("8080/tcp"),
	}

	//Starting the app Container
	customContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	//Stop tests if any errors encountered when setting up database connection
	if err != nil {
		return nil, err
	}

	return customContainer, nil
}

func getHostAndPorts(ctx context.Context, customContainer testcontainers.Container) (string, string, string, error) {
	//// Get the container's host and port
	cHost, err := customContainer.Host(ctx)
	if err != nil {
		return "", "", "", err
	}
	//obtaining the externally mapped port for the container
	cPort1, err := customContainer.MappedPort(ctx, "8080")
	if err != nil {
		return "", "", "", err
	}

	cPort2, err := customContainer.MappedPort(ctx, "9090")
	if err != nil {
		return "", "", "", err
	}

	return cHost, cPort1.Port(), cPort2.Port(), nil
}
