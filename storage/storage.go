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
	_, err := s.db.Exec(`
			CREATE TABLE IF NOT EXISTS counts (
			guild_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			counter INTEGER,
			PRIMARY KEY (guild_id, user_id)
			);`)
	if err != nil {
		panic(err)
	}
}
