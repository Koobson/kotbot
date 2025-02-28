package app

import (
	"fmt"
	"slices"
	"time"
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
		//TODO logging
		fmt.Println(err.Error())
		return
	}

	for _, guild := range guilds {
		fmt.Println(guild.Name)
		unverifiedRoleID, err := a.s.GetUnverifiedRoleID(guild.ID)
		if err != nil {
			//TODO logging
			fmt.Println(err.Error())
			continue
		}

		if unverifiedRoleID == "" {
			fmt.Println(guild.ID + " doesn't have it's unverified role setup")
			continue
		}

		guildMembers, err := a.dg.GuildMembers(guild.ID, "", 1000) //Rebuild if guild has >1000 members
		if err != nil {
			//TODO logging
			fmt.Println(err.Error())
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
			err = a.dg.GuildMemberDeleteWithReason(guild.ID, member.User.ID, "Kicked for not getting verified in time")
			if err != nil {
				//TODO logging
				continue
			}
		}
	}
}
