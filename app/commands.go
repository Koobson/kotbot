package app

import "github.com/bwmarrin/discordgo"

type commandHandler func(*discordgo.Session, *discordgo.InteractionCreate)

var commands map[string]discordgo.ApplicationCommand = map[string]discordgo.ApplicationCommand{
	"test": {Name: "test", Description: "testDescription"},
}

func (a *App) getCommandHandlers() map[string]commandHandler {
	return map[string]commandHandler{
		"test": a.commandTest,
	}
}
