package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Open the database connection
func Connect() (*sql.DB, error) {
	stringConnection := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
