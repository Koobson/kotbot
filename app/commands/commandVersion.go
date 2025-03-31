package commands

import (
	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandVersion(s *discordgo.Session, i *discordgo.InteractionCreate) {
	adminRoleID, err := cm.s.GetAdminRoleID(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandVersion()->cm.s.GetAdminRoleID(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandVersion()->s.Guild(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	if !utils.IsAdmin(g, i.Member, adminRoleID) {
		cm.InteractionRespond(s, i, utils.ErrorNoPermissions)
		logger.LogCommand("commandVersion()->!utils.IsAdmin(g, i.Member, adminRoleID)", nil, i.GuildID, i.Member.User.ID)
		return
	}

	versionString := "Kotbot ver. `" + cm.version + "`\nSource code: https://github.com/Koobson/kotbot"
	cm.InteractionRespond(s, i, versionString)

	logger.LogCommand("commandVersion()", nil, i.GuildID, i.Member.User.ID)
}
