export MIGRATION_DB="postgresql://golang_migrate:golang_migrate@localhost:5432/golang_migrate?sslmode=disable"

echo:
	echo $$MIGRATION_DB

command:
	go run cmd/command/main.go

api:
	go run cmd/command/main.go

test:
	go test -v ./...

test-coverage:
	go list ./... | grep -v '/cmd' | xargs go test -v -cover

create_migration:
	migrate create -ext=sql -dir=database/migrations -seq $(name)

migrate:
	migrate -path=database/migrations -database $$MIGRATION_DB -verbose up

rollback:
	migrate -path=database/migrations -database $$MIGRATION_DB -verbose down

.PHONY: test test-coverage
