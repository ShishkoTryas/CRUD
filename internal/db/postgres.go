package db

import (
	"database/sql"
)

func NewPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host= 127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=02012001")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
