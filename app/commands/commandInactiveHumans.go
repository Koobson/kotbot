package commands

import (
	"slices"
	"strings"
	"time"

	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
	"github.com/bwmarrin/discordgo"
)

const humanInactivityTime = time.Hour * 24 * 7

func (cm *CommandManager) commandInactiveHumans(s *discordgo.Session, i *discordgo.InteractionCreate) {
	adminRoleID, err := cm.s.GetAdminRoleID(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandInactiveHumans()->cm.s.GetAdminRoleID(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	g, err := s.Guild(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandInactiveHumans()->s.Guild(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}

	if !utils.IsAdmin(g, i.Member, adminRoleID) {
		cm.InteractionRespond(s, i, utils.ErrorNoPermissions)
		logger.LogCommand("commandInactiveHumans()->!utils.IsAdmin(g, i.Member, adminRoleID)", nil, i.GuildID, i.Member.User.ID)
		return
	}

	activeHumans, err := cm.s.GetActiveHumans(i.GuildID, time.Now().Add(-humanInactivityTime), 3)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandInactiveHumans()->cm.s.GetActiveHumans(i.GuildID, time.Now().Add(-humanInactivityTime), 3)", err, i.GuildID, i.Member.User.ID)
		return
	}

	guildMembers, err := s.GuildMembers(i.GuildID, "", 1000) //Rebuild if guild has >1000 members
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandInactiveHumans()->s.GuildMembers(i.GuildID, \"\", 1000)", err, i.GuildID, i.Member.User.ID)
		return
	}

	inactiveHumansBuilder := strings.Builder{}
	humanRoleID, err := cm.s.GetHumanRoleID(i.GuildID)
	if err != nil {
		cm.InteractionRespond(s, i, utils.ErrorInternal)
		logger.LogCommand("commandInactiveHumans()->cm.s.GetHumanRoleID(i.GuildID)", err, i.GuildID, i.Member.User.ID)
		return
	}
	if humanRoleID == "" {
		cm.InteractionRespond(s, i, utils.ErrorHumanRoleNotSer)
		logger.LogCommand("commandInactiveHumans()->humanRoleID == \"\"", nil, i.GuildID, i.Member.User.ID)
		return
	}

	for _, member := range guildMembers {
		if member.User.Bot {
			continue
		}
		if !slices.Contains(member.Roles, humanRoleID) {
			continue
		}
		if slices.ContainsFunc(activeHumans, func(activeHuman string) bool {
			return member.User.ID == activeHuman
		}) {
			continue
		}
		inactiveHumansBuilder.WriteString(member.User.Username + "\n")
	}

	cm.InteractionRespond(s, i, inactiveHumansBuilder.String())
	logger.LogCommand("commandInactiveHumans()", nil, i.GuildID, i.Member.User.ID)
}
