command:
	go run cmd/command/main.go

api:
	go run cmd/command/main.go

test:
	go test -v ./...

test-coverage:
	go list ./... | grep -v '/cmd' | xargs go test -v -cover

.PHONY: test test-coverage
