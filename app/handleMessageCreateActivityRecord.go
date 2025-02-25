package app

import (
	"github.com/bwmarrin/discordgo"
)

func (a *App) handleMessageCreateActivityRecord(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	a.s.AddHumansActivityTimestampRecord(m.GuildID, m.Author.ID)
}
