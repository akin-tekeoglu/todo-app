package utils

import (
	"database/sql"
)

func Db() *sql.DB {
	db, err := sql.Open("sqlite3", "tmp/dev.db")
	if err != nil {
		panic(err)
	}
	return db
}

func Execute(sql string, args ...interface{}) {
	db := Db()
	defer db.Close()
	query, err := db.Query(sql, args...)
	if err != nil {
		panic(err)
	}
	defer query.Close()
}

func Query(sql string, args ...interface{}) {
	// TODO use this library
	// https://github.com/jmoiron/sqlx
}
