package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)
	return db, err
}