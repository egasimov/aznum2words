###
test:
	go test ./...

app.generate:
	cd ./cmd/aznum2words-webapp/api \
	&& oapi-codegen -config models.cfg.yaml ./open-api-spec.yaml \
	&& oapi-codegen -config converter-server.cfg.yaml ./open-api-spec.yaml \
	&& oapi-codegen -config health-server.cfg.yaml ./open-api-spec.yaml

### WEBAPP
# we will put our integration testing in this path
WEBAPP_INTEGRATION_TEST_PATH?=./cmd/aznum2words-webapp/it

# set of env variables that you need for testing
WEBAPP_ENV_LOCAL_TEST=\
  APP_HOST=0.0.0.0 \
  APP_PORT=8080 \
  METRIC_PORT=9090 \
  DEPLOY_ENV=LOCAL

# this command will start a docker components that we set in docker-compose.webapp.integration.yml
webapp.docker.start.components:
	docker-compose -f ./docker-compose.webapp.integration.yml up -d

# shutting down docker components
webapp.docker.stop:
	docker-compose -f docker-compose.webapp.integration.yml down

# this command will trigger integration test
# WEBAPP_INTEGRATION_TEST_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./cmd/aznum2words-webapp/it directory
webapp.test.integration: webapp.docker.start.components
	$(WEBAPP_ENV_LOCAL_TEST) \
	go test -tags=integration $(WEBAPP_INTEGRATION_TEST_PATH) -count=1 -run=$(WEBAPP_INTEGRATION_TEST_PATH)

# this command will trigger integration test with verbose mode
webapp.test.integration.debug: webapp.docker.start.components
	go test -tags=integration $(WEBAPP_INTEGRATION_TEST_PATH) -count=1 -v -run=$(WEBAPP_INTEGRATION_TEST_PATH) \


#### CLI APP
# we will put our cli-app related integration testing in this path
CLIAPP_INTEGRATION_TEST_PATH?=./cmd/aznum2words-cli

cliapp.build:
	go build cmd/aznum2words-cli/aznum2words-cli.go

cliapp.build.integration:
	go build -tags integration cmd/aznum2words-cli/aznum2words-cli.go

cliapp.test.integration: cliapp.build.integration
	go test -tags=integration $(CLIAPP_INTEGRATION_TEST_PATH) -count=1 -run=$(CLIAPP_INTEGRATION_TEST_PATH)

cliapp.test.integration.debug: cliapp.build.integration
	go test -tags=integration $(CLIAPP_INTEGRATION_TEST_PATH) -count=1 -v -run=$(CLIAPP_INTEGRATION_TEST_PATH)
