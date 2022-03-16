dep:
	go mod download

build:
	go build ./...

test:
	@go test -coverpkg=./... -coverprofile cov.out  $(shell go list ./... | grep -v /vendor/)
	@go tool cover -func=cov.out

install:
	go install ./...

clean:
	go clean ./...

analyze: dep
	go fmt ./...
	gosec ./...
	staticcheck ./...