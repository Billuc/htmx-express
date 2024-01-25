package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const DB_NAME string = "data.db"

var migrations = [2]string{CREATE_TAGS, CREATE_CARDS}

var db *sql.DB

func GetDB() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	db, err = sql.Open("sqlite3", DB_NAME)

	if err != nil {
		return nil, err
	}

	return migrate(db)
}

func migrate(db *sql.DB) (*sql.DB, error) {
	for _, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
