package app

import (
	"github.com/bwmarrin/discordgo"
)

func (a *App) handleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	handler, ok := a.getCommandHandlers()[i.ApplicationCommandData().Name]
	if !ok {
		panic("command not found")
	}
	handler(s, i)
}
