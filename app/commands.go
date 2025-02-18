package app

//TODO add command to set human role

import "github.com/bwmarrin/discordgo"

type commandHandler func(*discordgo.Session, *discordgo.InteractionCreate)

var commands map[string]discordgo.ApplicationCommand = map[string]discordgo.ApplicationCommand{
	"test":                {Name: "test", Description: "testDescription"},
	"get_inactive_humans": {Name: "get_inactive_humans", Description: "gives inactive humans"},
}

func (a *App) getCommandHandlers() map[string]commandHandler {
	return map[string]commandHandler{
		"test":                a.commandTest,
		"get_inactive_humans": a.commandInactiveHumans,
	}
}
