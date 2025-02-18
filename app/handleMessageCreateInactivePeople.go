package app

import (
	"github.com/bwmarrin/discordgo"
)

func (a *App) handleMessageCreateInactiveHumans(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	a.s.AddInactiveHumansRecord(m.GuildID, m.Author.ID)
}
