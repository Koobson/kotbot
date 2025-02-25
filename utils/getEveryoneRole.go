package utils

import "github.com/bwmarrin/discordgo"

func GetEveryoneRole(g discordgo.Guild) *discordgo.Role {
	roles := g.Roles

	for _, role := range roles {
		if role.Name == "@everyone" {
			return role
		}
	}
	return nil
}
