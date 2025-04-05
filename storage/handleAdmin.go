package storage

import (
	"database/sql"
)

func (s *Storage) initAdmin() {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	_, err := s.db.Exec(`
			CREATE TABLE IF NOT EXISTS admin_role (
			guild_id TEXT NOT NULL,
			role_id TEXT NOT NULL,
			PRIMARY KEY (guild_id)
			);`)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) SetAdminRoleID(guildID string, adminRoleID string) error {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	_, err := s.db.Exec(`INSERT INTO admin_role(guild_id, role_id) VALUES($1,$2)
  						 ON CONFLICT(guild_id) DO UPDATE SET role_id=excluded.role_id`, guildID, adminRoleID)
	return err
}

func (s *Storage) GetAdminRoleID(guildID string) (string, error) {
	s.dbLock.Lock()
	defer s.dbLock.Unlock()
	adminRoleID := ""
	err := s.db.Get(&adminRoleID, "SELECT role_id FROM admin_role WHERE guild_id=$1", guildID)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return adminRoleID, nil
}
