command:
	go run cmd/command/main.go

api:
	go run cmd/command/main.go

test:
	go test -v ./...

test-coverage:
	go list ./... | grep -v '/cmd' | xargs go test -v -cover

migrate:
	migrate -path=database/migrations -database "postgresql://root:root@localhost:5432/expenses?sslmode=disable" -verbose up

rollback:
	migrate -path=database/migrations -database "postgresql://root:root@localhost:5432/expenses?sslmode=disable" -verbose down

.PHONY: test test-coverage
