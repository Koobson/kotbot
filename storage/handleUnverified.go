package storage

import (
	"database/sql"
)

func (s *Storage) initUnverified() {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	_, err := s.db.Exec(`
			CREATE TABLE IF NOT EXISTS unverified_role (
			guild_id TEXT NOT NULL,
			role_id TEXT NOT NULL,
			PRIMARY KEY (guild_id)
			);`)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) SetUnverifiedRoleID(guildID string, unverifiedRoleID string) error {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	_, err := s.db.Exec(`INSERT INTO unverified_role(guild_id, role_id) VALUES($1,$2)
  						 ON CONFLICT(guild_id) DO UPDATE SET role_id=excluded.role_id`, guildID, unverifiedRoleID)
	return err
}

func (s *Storage) GetUnverifiedRoleID(guildID string) (string, error) {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	unverifiedRoleID := ""
	err := s.db.Get(&unverifiedRoleID, "SELECT role_id FROM unverified_role WHERE guild_id=$1", guildID)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return unverifiedRoleID, nil
}
