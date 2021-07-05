package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Db() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", os.Getenv("DB"))
	if err != nil {
		panic(err)
	}
	return db
}
