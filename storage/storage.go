package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sqlx.DB
}

func New(dbFilename string) *Storage {
	db, err := sqlx.Open("sqlite3", dbFilename)
	if err != nil {
		panic(err)
	}

	s := Storage{db: db}
	s.init()

	return &s
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) init() {
	s.initInactiveHumans()
}
