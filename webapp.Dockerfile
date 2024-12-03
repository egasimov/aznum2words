# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /aznum2words-go-svc -v ./cmd/aznum2words-webapp/aznum2words-webapp.go

##
## Deploy
##
#FROM golang:1.19.1-buster
#FROM gcr.io/distroless/base-debian10
FROM scratch

WORKDIR /

COPY --from=build /aznum2words-go-svc /app/aznum2words-go-svc-exec

EXPOSE 8080
EXPOSE 9090

#USER root

WORKDIR /app

ENTRYPOINT ["./aznum2words-go-svc-exec"]