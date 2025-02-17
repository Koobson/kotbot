package storage

type counterRow struct {
	GuildID string `db:"guild_id"`
	UserID  string `db:"user_id"`
	Counter int    `db:"counter"`
}

func (s *Storage) GetCounter(guildID string) map[string]int {
	rows := []counterRow{}
	err := s.db.Select(&rows, "SELECT * FROM counts WHERE guild_id=$1", guildID)

	if err != nil {
		panic(err)
	}

	retCounts := map[string]int{}
	for _, v := range rows {
		retCounts[v.UserID] = v.Counter
	}
	return retCounts
}

func (s *Storage) CounterIncrement(guildID string, userID string) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	row := counterRow{}
	err = tx.Get(&row, "SELECT * FROM counts where guild_id=$1 AND user_id=$2", guildID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("UPDATE counts SET counter=$1 WHERE guild_id=$2 AND user_id=$3", row.Counter+1, guildID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
