package commands

import (
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	helpString :=
		"`/help` Displays all commands.\n\n" +
			"`/version` Displays bot version. \n\n" +
			"`/get_inactive_humans` Gives a list of people with a Human role that have been inactive for a week.\n\n" +
			"`/set_human_role` Sets the Human role.\n\n" +
			"`/set_admin_role` Sets the Admin role that grants access to other commands. This command can only be used by the server owner.\n\n" +
			"`/set_unverified_role` Sets the Unverified role. People with this role will be automatically kicked after a week."
	cm.InteractionRespond(s, i, helpString)

	logger.LogCommand("commandHelp()", nil, i.GuildID, i.Member.User.ID)
}
