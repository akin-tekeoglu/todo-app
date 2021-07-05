package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	m, err := migrate.New(
		"file://./migrations",
		os.Getenv("DB"))
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
