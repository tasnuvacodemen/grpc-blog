package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(dbstring string) (*Storage,error) {
	db, err := sqlx.Connect("postgres", dbstring)
	if err !=nil {
		return nil,err
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return &Storage{db:db},nil
}


func NewTestStorage(dbstring string, migrationDir string) (*Storage, func()) {
	db, teardown := MustNewDevelopmentDB(dbstring, migrationDir)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return &Storage{db: db}, teardown
}
