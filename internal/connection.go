package internal

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1)/dbname")
	if err != nil {
		return nil, err
	}

	return db, nil
}
