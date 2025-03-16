package commands

import (
	"github.com/Koobson/kotbot/storage"
	"github.com/bwmarrin/discordgo"
)

type CommandManager struct {
	dg      *discordgo.Session
	s       *storage.Storage
	version string
}

func New(dg *discordgo.Session, storage *storage.Storage, version string) *CommandManager {
	cm := CommandManager{dg: dg, s: storage, version: version}
	return &cm
}

func (cm *CommandManager) InteractionRespond(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	s.InteractionRespond(i.Interaction,
		&discordgo.InteractionResponse{Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: message}})
}

func (cm *CommandManager) GetCommands() map[string]discordgo.ApplicationCommand {
	return commands
}
