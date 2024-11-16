build:
	@go build -o bin/gofeed cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/gofeed

migration:
	@migrate create -ext sql -dir db/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run db/migrate/main.go up

migrate-down:
	@go run db/migrate/main.go down
