package commands

import (
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandVersion(s *discordgo.Session, i *discordgo.InteractionCreate) {
	versionString := "Kotbot ver. `" + cm.version + "`\nSource code: https://github.com/Koobson/kotbot"
	cm.InteractionRespond(s, i, versionString)

	logger.LogCommand("commandHelp()", nil, i.GuildID, i.Member.User.ID)
}
