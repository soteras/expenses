export MIGRATION_DB="postgresql://root:root@localhost:5432/expenses?sslmode=disable"
export MIGRATION_PATH="database/migrations"

command:
	GO_ENV=development go run cmd/command/main.go

api:
	GO_ENV=development go run cmd/api/main.go

test:
	GO_ENV=test go test -v ./...

coverage:
	GO_ENV=test go list ./... | grep -v '/cmd' | xargs go test -v -cover

create_migration:
	migrate create -ext=sql -dir=$(MIGRATION_PATH) -seq $(name)

migrate:
	migrate -path=$(MIGRATION_PATH) -database $(MIGRATION_DB) -verbose up

rollback:
	migrate -path=$(MIGRATION_PATH) -database $(MIGRATION_DB) -verbose down
