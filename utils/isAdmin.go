package utils

import (
	"slices"

	"github.com/bwmarrin/discordgo"
)

func IsAdmin(g *discordgo.Guild, m *discordgo.Member, adminRoleID string) bool {
	if adminRoleID == GetEveryoneRole(*g).ID {
		return true
	}
	if slices.Contains(m.Roles, adminRoleID) {
		return true
	}
	return false
}
