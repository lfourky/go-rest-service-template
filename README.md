# Go REST service template

Template for REST APIs written in Go.

---

## Working with database migrations

Create `go_rest_service` database in your PostgreSQL instance

Install **sql-migrate**

    go get -v github.com/rubenv/sql-migrate/...

Position yourself in the **db_migrations** directory and run

    sql-migrate up --env local

## Working with unit tests

Position yourself in the **root** directory and run the tests

    go test ./...

## Working with integration tests

Create `go_rest_service` database in your PostgreSQL instance

(Optional) Position yourself in the **migrations/** directory and run

    sql-migrate up --env integration_test

Position yourself in the **root** directory and run the tests

    go test ./... -tags integration

## Working with the linter

Install **golangci-lint**

    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

Position yourself in the root directory and run

    golangci-lint run --allow-parallel-runners --enable-all -e dupl --timeout 30m
