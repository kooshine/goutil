package db

import (
	"database/sql"
	"fmt"
)

func OpenDB(dbUrl string) (*sql.DB) {
	dbpool, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	return dbpool
}
