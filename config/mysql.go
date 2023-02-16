package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/data_user")
	if err != nil {
		return nil, err
	}

	return db, nil
}
