package app

import (
	"fmt"
	"slices"
	"time"

	"github.com/Koobson/kotbot/utils/logger"
)

const inactivityKickAfterTime = time.Hour * 24 * 7
const jobInterval = time.Hour

func (a *App) startJobKickUnverified() {
	go func() {
		a.kickUnverified()
		ticker := time.NewTicker(jobInterval)
		for range ticker.C {
			a.kickUnverified()
		}

	}()
}

func (a *App) kickUnverified() {
	guilds, err := a.dg.UserGuilds(200, "", "", false) //Rebuild if bot on > 200 guilds
	if err != nil {
		logger.Log("kickUnverified()->a.dg.UserGuilds(200, \"\", \"\", false)", err)
		return
	}

	kickedUserCounter := 0
	for _, guild := range guilds {
		unverifiedRoleID, err := a.s.GetUnverifiedRoleID(guild.ID)
		if err != nil {
			logger.Log("kickUnverified()->a.s.GetUnverifiedRoleID(guild.ID)(200, \"\", \"\", false)", err, logger.GuildID(guild.ID))
			continue
		}

		if unverifiedRoleID == "" {
			logger.Log("kickUnverified()->unverifiedRoleID == \"\"", fmt.Errorf("%s doesn't have it's unverified role setup", guild.ID),
				logger.GuildID(guild.ID))
			continue
		}

		guildMembers, err := a.dg.GuildMembers(guild.ID, "", 1000) //Rebuild if guild has >1000 members
		if err != nil {
			logger.Log("kickUnverified()->a.dg.GuildMembers(guild.ID, \"\", 1000)", err, logger.GuildID(guild.ID))
			continue
		}

		for _, member := range guildMembers {
			if member.User.Bot {
				continue
			}

			if !slices.Contains(member.Roles, unverifiedRoleID) {
				continue
			}
			if member.JoinedAt.After(time.Now().Add(-inactivityKickAfterTime)) {
				continue
			}
			err = a.dg.GuildMemberDeleteWithReason(guild.ID, member.User.ID, "Wyrzucono za brak weryfikacji")
			if err != nil {
				logger.Log("kickUnverified()->a.dg.GuildMemberDeleteWithReason(guild.ID, member.User.ID, \"Wyrzucono za brak weryfikacji\")", err, logger.GuildID(guild.ID), logger.UserID(member.User.ID))
				continue
			}
			kickedUserCounter++
		}
	}

	logger.Log("kickUnverified()", nil, logger.GuildCount(len(guilds)), logger.KickedUserCount(kickedUserCounter))
}
