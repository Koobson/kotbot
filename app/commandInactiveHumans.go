package app

import (
	"slices"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (a *App) commandInactiveHumans(s *discordgo.Session, i *discordgo.InteractionCreate) {
	activeHumans, err := a.s.GetActiveHumans(i.GuildID, time.Now().Add(-time.Hour*24*7), 3)
	if err != nil {
		a.InteractionRespond(s, i, err.Error())
		return
	}

	guildMembers, err := s.GuildMembers(i.GuildID, "", 1000) //Rebuild if guild has >1000 members
	if err != nil {
		a.InteractionRespond(s, i, err.Error())
		return
	}

	inactiveHumansBuilder := strings.Builder{}
	humanRoleID, err := a.s.GetHumanRoleID(i.GuildID)
	if err != nil {
		a.InteractionRespond(s, i, err.Error())
		return
	}
	if humanRoleID == "" {
		a.InteractionRespond(s, i, "Human role is not set")
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

	a.InteractionRespond(s, i, inactiveHumansBuilder.String())
}
