package commands

import (
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandSetAdminRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, err.Error())
		return
	}

	if i.Member.User.ID != g.OwnerID {
		cm.InteractionRespond(s, i, "Only the owner can invoke this command")
		return
	}

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		cm.InteractionRespond(s, i, "Recieved no options in set_admin_role command")
		return
	}

	adminRole := options[0].RoleValue(s, i.GuildID)
	cm.s.SetAdminRoleID(i.GuildID, adminRole.ID)
	cm.InteractionRespond(s, i, adminRole.Mention()+" is now an admin role")
}
