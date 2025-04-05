package storage

import (
	"database/sql"
	"time"
)

type activeHumansRow struct {
	GuildID   string `db:"guild_id"`
	UserID    string `db:"user_id"`
	Timestamp string `db:"timestamp"`
}

func (s *Storage) initHuman() {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	_, err := s.db.Exec(`
			CREATE TABLE IF NOT EXISTS message_timestamps (
			guild_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			timestamp TEXT NOT NULL,
			PRIMARY KEY (guild_id, user_id, timestamp)
			);`)
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(`
			CREATE TABLE IF NOT EXISTS human_role (
			guild_id TEXT NOT NULL,
			role_id TEXT NOT NULL,
			PRIMARY KEY (guild_id)
			);`)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) SetHumanRoleID(guildID string, humanRoleID string) error {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	_, err := s.db.Exec(`INSERT INTO human_role(guild_id, role_id) VALUES($1,$2)
  						 ON CONFLICT(guild_id) DO UPDATE SET role_id=excluded.role_id`, guildID, humanRoleID)
	return err
}

func (s *Storage) GetHumanRoleID(guildID string) (string, error) {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	humanRoleID := ""
	err := s.db.Get(&humanRoleID, "SELECT role_id FROM human_role WHERE guild_id=$1", guildID)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return humanRoleID, nil
}

func (s *Storage) GetActiveHumans(guildID string, cutoffDate time.Time, cutoffMessagesCount int) ([]string, error) {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	rows := []activeHumansRow{}
	err := s.db.Select(&rows, "SELECT * FROM message_timestamps WHERE guild_id=$1", guildID)

	if err != nil {
		return nil, err
	}

	counts := map[string]int{}
	for _, v := range rows {
		currentMessageTime, err := time.Parse(time.RFC3339Nano, v.Timestamp)
		if err != nil {
			return nil, err
		}
		if currentMessageTime.After(cutoffDate) {
			counts[v.UserID]++
		}
	}

	activeHumans := []string{}
	for userID, messageCount := range counts {
		if messageCount >= cutoffMessagesCount {
			activeHumans = append(activeHumans, userID)
		}
	}

	return activeHumans, nil
}

func (s *Storage) AddHumansActivityTimestampRecord(guildID string, userID string) error {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	_, err := s.db.Exec("INSERT INTO message_timestamps VALUES ($1, $2, $3)", guildID, userID, time.Now().Format(time.RFC3339Nano))
	if err != nil {
		return err
	}

	return nil
}
