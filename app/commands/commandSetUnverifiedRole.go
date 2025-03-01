package commands

import (
	"errors"

	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandSetUnverifiedRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	adminRoleID, err := cm.s.GetAdminRoleID(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetUnverifiedRole()->cm.s.GetAdminRoleID(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetUnverifiedRole()->s.Guild(i.GuildID))", err, i.GuildID, i.Member.User.ID)
		return
	}

	if !utils.IsAdmin(g, i.Member, adminRoleID) {
		cm.InteractionRespond(s, i, utils.ErrorNoPermissions)
		logger.LogCommand("commandSetUnverifiedRole()->!utils.IsAdmin(g, i.Member, adminRoleID)", nil, i.GuildID, i.Member.User.ID)
		return
	}

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetUnverifiedRole()->len(options) == 0", errors.New("discordgo.ApplicationCommandInteractionOption does not exist"),
			i.GuildID, i.Member.User.ID)
		return
	}

	unverifiedRole := options[0].RoleValue(s, i.GuildID)
	cm.s.SetUnverifiedRoleID(i.GuildID, unverifiedRole.ID)
	cm.InteractionRespond(s, i, unverifiedRole.Mention()+" is now an unverified role")
}
