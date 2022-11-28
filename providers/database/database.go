package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const connectionString = "root:admin2022@tcp(localhost:3306)/customers_db?parseTime=true"

func connect() *sql.DB {
	cn, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	return cn
}

func ExecuteQuery(query string, data ...interface{}) *sql.Result {

	db := connect()
	defer (func() {
		db.Close()
	})()

	result, err := db.Exec(query, data...)
	if err != nil {
		panic(err)
	}

	return &result
}

func GetData(query string, data ...interface{}) *sql.Rows {
	db := connect()
	defer (func() {
		db.Close()
	})()

	rows, err := db.Query(query, data...)
	if err != nil {
		panic(err)
	}

	return rows
}
