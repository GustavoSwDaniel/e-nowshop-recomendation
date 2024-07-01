package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectionDatabase(DatabaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}
