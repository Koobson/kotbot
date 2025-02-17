package app

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (a *App) handleMessageCreatePingPong(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(strings.ToLower(m.Content), "meow") {
		err := a.s.CounterIncrement(m.GuildID, m.Author.ID)
		if err != nil {
			panic(err)
		}
		s.ChannelMessageSend(m.ChannelID, "meow :3")
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
