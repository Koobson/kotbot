package app

import (
	"fmt"

	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (a *App) handleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	handler, ok := a.cm.GetCommandHandlers()[i.ApplicationCommandData().Name]
	if !ok {
		logger.Log("handleCommands()->a.cm.GetCommandHandlers()[i.ApplicationCommandData().Name]", fmt.Errorf("command not found: %s", i.ApplicationCommandData().Name))
		panic("command not found")
	}
	handler(s, i)
}
