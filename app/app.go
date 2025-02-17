package app

import (
	"github.com/bwmarrin/discordgo"
)

type App struct {
	dg *discordgo.Session
}

func New(botToken string) *App {
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		panic(err)
	}
	dg.Identify.Intents = discordgo.IntentsAll

	a := App{dg: dg}
	a.registerHandlers()

	return &a
}

func (a *App) registerHandlers() {
	a.dg.AddHandler(a.handleGuildCreate)
	a.dg.AddHandler(a.handleMessageCreatePingPong)
	a.dg.AddHandler(a.handleCommands)
}

func (a *App) Start() {
	err := a.dg.Open()
	if err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	a.dg.Close()
}
