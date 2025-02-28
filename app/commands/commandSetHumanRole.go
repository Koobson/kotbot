package commands

import (
	"github.com/Koobson/kotbot/utils"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandSetHumanRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	adminRoleID, err := cm.s.GetAdminRoleID(i.GuildID) //Segfault here
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
	humanRole := options[0].RoleValue(s, i.GuildID)

	cm.s.SetHumanRoleID(i.GuildID, humanRole.ID)
	cm.InteractionRespond(s, i, humanRole.Mention()+" is now a human role")
}
