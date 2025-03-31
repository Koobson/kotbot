package commands

import (
	"errors"

	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandSetHumanRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetHumanRole()->s.Guild(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	if i.Member.User.ID != g.OwnerID {
		cm.InteractionRespond(s, i, utils.ErrorNoPermissions)
		logger.LogCommand("commandSetHuman()->i.Member.User.ID != g.OwnerID", err, i.GuildID, i.Member.User.ID)
		return
	}

	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandSetHumanRole()->len(options) == 0", errors.New("discordgo.ApplicationCommandInteractionOption does not exist"),
			i.GuildID, i.Member.User.ID)
		return
	}

	humanRole := options[0].RoleValue(s, i.GuildID)
	cm.s.SetHumanRoleID(i.GuildID, humanRole.ID)
	cm.InteractionRespond(s, i, humanRole.Mention()+" są teraz dwunożnymi")
	logger.LogCommand("commandSetHumanRole()", nil, i.GuildID, i.Member.User.ID)
}
