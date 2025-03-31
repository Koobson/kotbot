package commands

import (
	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	adminRoleID, err := cm.s.GetAdminRoleID(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandHelp()->cm.s.GetAdminRoleID(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandHelp()->s.Guild(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	if !utils.IsAdmin(g, i.Member, adminRoleID) {
		cm.InteractionRespond(s, i, utils.ErrorNoPermissions)
		logger.LogCommand("commandHelp()->!utils.IsAdmin(g, i.Member, adminRoleID)", nil, i.GuildID, i.Member.User.ID)
		return
	}

	helpString :=
		"`/help` Displays all commands.\n\n" +
			"`/version` Displays bot version.\n\n" +
			"`/get_inactive_humans` Gives a list of people with a `Human` role that have been inactive for a week.\n\n" +
			"`/set_human_role` Sets the `Human` role. This command can only be used by the server owner.\n\n" +
			"`/set_admin_role` Sets the `Admin` role that grants access to other commands (those not restricted by owner status). This command can only be used by the server owner.\n\n" +
			"`/set_unverified_role` Sets the `Unverified` role. People with this role will be automatically kicked after a week. This command can only be used by the server owner.\n\n" +
			"`/stats` Displays some statistics collected during the bot's lifetime."
	cm.InteractionRespond(s, i, helpString)

	logger.LogCommand("commandHelp()", nil, i.GuildID, i.Member.User.ID)
}
