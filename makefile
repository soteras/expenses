export MIGRATION_DB="postgresql://root:root@localhost:5432/expenses?sslmode=disable"
export MIGRATION_PATH="database/migrations"

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
	migrate create -ext=sql -dir=$(MIGRATION_PATH) -seq $(name)

migrate:
	migrate -path=$(MIGRATION_PATH) -database $(MIGRATION_DB) -verbose up

rollback:
	migrate -path=$(MIGRATION_PATH) -database $(MIGRATION_DB) -verbose down

.PHONY: test test-coverage
