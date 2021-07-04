package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://./migrations",
		"sqlite3://./storage/dev.db")
	if err != nil {
		panic(err)
	}
	err = m.Up()
	// It throws an exception when no migrations applied
	// And there is no way to turn it off
	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}
