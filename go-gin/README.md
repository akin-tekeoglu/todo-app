# Requirements

- Go 1.16

# Installation

- `go install`
- `mkdir -p storage && sqlite3 storage/dev.db "create table temp(field int); drop table temp;"`
- `go run ./migrations/migrate.go`
- `go run main.go`