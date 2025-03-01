package commands

import (
	"errors"

	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandSetAdminRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetAdminRole()->s.Guild(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	if i.Member.User.ID != g.OwnerID {
		cm.InteractionRespond(s, i, utils.ErrorNoPermissions)
		logger.LogCommand("commandSetAdminRole()->i.Member.User.ID != g.OwnerID", err, i.GuildID, i.Member.User.ID)
		return
	}

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetAdminRole()->len(options) == 0", errors.New("discordgo.ApplicationCommandInteractionOption does not exist"),
			i.GuildID, i.Member.User.ID)
		return
	}

	adminRole := options[0].RoleValue(s, i.GuildID)
	cm.s.SetAdminRoleID(i.GuildID, adminRole.ID)
	cm.InteractionRespond(s, i, adminRole.Mention()+" ma teraz uprawnienia administratora")
	logger.LogCommand("commandSetAdminRole()", nil, i.GuildID, i.Member.User.ID)
}
