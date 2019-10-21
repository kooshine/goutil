package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func OpenDB(dbUrl string) (*sql.DB) {
	dbpool, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	return dbpool
}
