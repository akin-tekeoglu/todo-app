package utils

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Db() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
