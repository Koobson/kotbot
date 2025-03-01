package app

import (
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (a *App) handleGuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, g.ID)
	if err != nil {
		logger.Log("handleGuildCreate()->s.ApplicationCommands(s.State.User.ID, g.ID)", err, logger.GuildID(g.ID))
		return
	}

	for _, cmd := range registeredCommands {
		err := a.dg.ApplicationCommandDelete(s.State.User.ID, g.ID, cmd.ID)
		if err != nil {
			logger.Log("handleGuildCreate()->a.dg.ApplicationCommandDelete(s.State.User.ID, g.ID, cmd.ID)", err, logger.GuildID(g.ID))
			return
		}
	}

	for _, command := range a.cm.GetCommands() {
		_, err := a.dg.ApplicationCommandCreate(s.State.User.ID, g.ID, &command)
		if err != nil {
			logger.Log("handleGuildCreate()->a.dg.ApplicationCommandCreate(s.State.User.ID, g.ID, &command)", err, logger.GuildID(g.ID))
			return
		}
	}

}
