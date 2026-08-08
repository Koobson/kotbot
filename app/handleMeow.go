package app

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var meowList = []string{
	"meow :3",
	"meow =w=",
	"mrreow :3",
	"mew >w<",
	"meoww ^-^",
	"mrow :3",
	"nyaa~",
	"purrmeow ~",
	"mrrp :3",
}

func (a *App) handleMeow(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(strings.ToLower(m.Content), "meow") {
		meowResponse := meowList[rand.Intn(len(meowList))]
		s.ChannelMessageSend(m.ChannelID, meowResponse)
	}
}
