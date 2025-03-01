package app

import (
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (a *App) handleMessageCreateActivityRecord(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	err := a.s.AddHumansActivityTimestampRecord(m.GuildID, m.Author.ID)
	if err != nil {
		logger.Log("handleMessageCreateActivityRecord()->a.s.AddHumansActivityTimestampRecord(m.GuildID, m.Author.ID)",
			err, logger.GuildID(m.GuildID), logger.UserID(m.Author.ID))
	}
}
