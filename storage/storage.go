package storage

import (
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sqlx.DB
}

func New(dbFilename string) *Storage {
	db, err := sqlx.Open("sqlite", dbFilename)
	if err != nil {
		logger.Log("New()->sqlx.Open(\"sqlite3\", dbFilename)", err)
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
	s.initHuman()
	s.initAdmin()
	s.initUnverified()
	s.initKickStats()
}
