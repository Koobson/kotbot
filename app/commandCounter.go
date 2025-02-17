package app

import (
	"slices"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (a *App) commandCounter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	counter := a.s.GetCounter(i.GuildID)
	keys := make([]string, 0, len(counter))

	for key := range counter {
		keys = append(keys, key)
	}

	slices.SortStableFunc(keys, func(i, j string) int {
		return counter[j] - counter[i]
	})

	builder := strings.Builder{}

	for _, v := range keys {
		user, err := s.User(v)
		if err != nil {
			panic(err)
		}
		builder.WriteString(user.GlobalName + ": " + strconv.Itoa(counter[v]) + "\n")
	}

	s.InteractionRespond(i.Interaction,
		&discordgo.InteractionResponse{Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: builder.String()}})
}
