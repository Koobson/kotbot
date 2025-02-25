package app

import (
	"github.com/bwmarrin/discordgo"
)

func (a *App) handleGuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, g.ID)
	if err != nil {
		panic(err)
	}

	for _, cmd := range registeredCommands {
		err := a.dg.ApplicationCommandDelete(s.State.User.ID, g.ID, cmd.ID)
		if err != nil {
			panic(err)
		}
	}

	for _, command := range a.cm.GetCommands() {
		_, err := a.dg.ApplicationCommandCreate(s.State.User.ID, g.ID, &command)
		if err != nil {
			panic(err)
		}
	}

}
