package commands

import (
	"errors"

	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandSetUnverifiedRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetUnverifiedRole()->s.Guild(i.GuildID))", err, i.GuildID, i.Member.User.ID)
		return
	}

	if i.Member.User.ID != g.OwnerID {
		cm.InteractionRespond(s, i, utils.ErrorNoPermissions)
		logger.LogCommand("commandSetUnverifiedRole()->i.Member.User.ID != g.OwnerID", err, i.GuildID, i.Member.User.ID)
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
	cm.InteractionRespond(s, i, unverifiedRole.Mention()+" są uznawani za niezweryfikowanych")
	logger.LogCommand("commandSetUnverifiedRole()", nil, i.GuildID, i.Member.User.ID)
}
