package commands

import (
	"strconv"

	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

func (cm *CommandManager) commandStats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	unverifiedUsersCount, err := cm.s.GetKickedUsersCount(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandStats()->cm.s.GetKickedUsersCount(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	cm.InteractionRespond(s, i, "Do tej pory wyrzucono `"+strconv.Itoa(unverifiedUsersCount)+"` Użytkowników")
	logger.LogCommand("commandStats()", nil, i.GuildID, i.Member.User.ID)
}
