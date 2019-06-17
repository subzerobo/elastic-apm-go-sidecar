GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=groot_osx
BINARY_UNIX=groot
VERSION=2.0.0

all: build
proto:
	protoc --go_out=. ./protos/*.proto
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	ELASTIC_APM_SERVER_URL=http://yourserver.com ELASTIC_APM_SERVER_TIMEOUT=30s ./$(BINARY_NAME) start
debug:
	$(GORUN) main.go config.go ndjson.go start -listen :1199
deps:
	$(GOGET)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v -ldflags "-X main.VERSION=$(VERSION)"
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
