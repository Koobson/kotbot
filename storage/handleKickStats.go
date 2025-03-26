package storage

import (
	"database/sql"
	"strconv"
	"time"
)

func (s *Storage) initKickStats() {
	_, err := s.db.Exec(`
			CREATE TABLE IF NOT EXISTS kick_stats (
			guild_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			timestamp TEXT NOT NULL,
			PRIMARY KEY (guild_id, user_id, timestamp)
			);`)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) AddKickedUser(guildID string, userID string) error {
	_, err := s.db.Exec(`INSERT INTO kick_stats(guild_id, user_id, timestamp) VALUES($1,$2,$3)`, guildID, userID, time.Now().Format(time.RFC3339Nano))
	return err
}

func (s *Storage) GetKickedUsersCount(guildID string) (int, error) {
	kickedUsersCountString := ""
	err := s.db.Get(&kickedUsersCountString, "SELECT COUNT(*) FROM kick_stats WHERE guild_id=$1", guildID)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, nil
		}
		return -1, err
	}
	kickedUsersCount, err := strconv.Atoi(kickedUsersCountString)
	if err != nil {
		return -1, nil
	}
	return kickedUsersCount, nil
}
