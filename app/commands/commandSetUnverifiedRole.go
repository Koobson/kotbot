package commands

import (
	"github.com/Koobson/kotbot/utils"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandSetUnverifiedRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	adminRoleID, err := cm.s.GetAdminRoleID(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, err.Error())
		return
	}

	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, err.Error())
		return
	}

	if !utils.IsAdmin(g, i.Member, adminRoleID) {
		cm.InteractionRespond(s, i, "you do not have permissions to use this command")
		return
	}

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		cm.InteractionRespond(s, i, "Received no options in set_unverified_role command")
		return
	}

	unverifiedRole := options[0].RoleValue(s, i.GuildID)
	cm.s.SetUnverifiedRoleID(i.GuildID, unverifiedRole.ID)
	cm.InteractionRespond(s, i, unverifiedRole.Mention()+" is now an unverified role")
}
