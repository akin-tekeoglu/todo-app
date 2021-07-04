package utils

import (
	"database/sql"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

func Db() *sql.DB {
	db, err := sql.Open("sqlite3", "file:storage/dev.db")
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

func Query(sql string, result interface{}, extractor func(rows *sql.Rows) interface{}, args ...interface{}) interface{} {
	db := Db()
	defer db.Close()
	rows, err := db.Query(sql, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	elemSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(result)), 0, 10)
	for rows.Next() {
		object := extractor(rows)
		elemSlice = reflect.Append(elemSlice, reflect.ValueOf(object))
	}
	return elemSlice.Interface()
}

func Single(sql string, result interface{}, extractor func(rows *sql.Rows) interface{}, args ...interface{}) interface{} {
	list := Query(sql, result, extractor, args...)
	return reflect.ValueOf(list).Index(0)
}
