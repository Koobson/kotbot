package commands

import "github.com/bwmarrin/discordgo"

type commandHandler func(*discordgo.Session, *discordgo.InteractionCreate)

var commands map[string]discordgo.ApplicationCommand = map[string]discordgo.ApplicationCommand{
	"get_inactive_humans": {Name: "get_inactive_humans", Description: "gives inactive humans"},
	"set_human_role": {Name: "set_human_role", Description: "sets human role", Options: []*discordgo.ApplicationCommandOption{
		{Type: discordgo.ApplicationCommandOptionRole, Name: "human_role", Description: "human role", Required: true}}},
	"set_admin_role": {Name: "set_admin_role", Description: "sets admin role", Options: []*discordgo.ApplicationCommandOption{
		{Type: discordgo.ApplicationCommandOptionRole, Name: "admin_role", Description: "admin role", Required: true}}},
	"set_unverified_role": {Name: "set_unverified_role", Description: "sets unverified role", Options: []*discordgo.ApplicationCommandOption{
		{Type: discordgo.ApplicationCommandOptionRole, Name: "unverified_role", Description: "unverified role", Required: true}}},
	"help":    {Name: "help", Description: "displays available commands"},
	"version": {Name: "version", Description: "displays bot version"},
}

func (cm *CommandManager) GetCommandHandlers() map[string]commandHandler {
	return map[string]commandHandler{
		"get_inactive_humans": cm.commandInactiveHumans,
		"set_human_role":      cm.commandSetHumanRole,
		"set_admin_role":      cm.commandSetAdminRole,
		"set_unverified_role": cm.commandSetUnverifiedRole,
		"help":                cm.commandHelp,
		"version":             cm.commandVersion,
	}
}
