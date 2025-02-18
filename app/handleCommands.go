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

func (a *App) InteractionRespond(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	s.InteractionRespond(i.Interaction,
		&discordgo.InteractionResponse{Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: message}})
}
