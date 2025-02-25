package app

import (
	"github.com/Koobson/kotbot/app/commands"
	"github.com/Koobson/kotbot/storage"
	"github.com/bwmarrin/discordgo"
)

type App struct {
	dg *discordgo.Session
	s  *storage.Storage
	cm *commands.CommandManager
}

func New(botToken string, storage *storage.Storage) *App {
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		panic(err)
	}
	dg.Identify.Intents = discordgo.IntentsAll

	a := App{dg: dg, s: storage}
	a.registerHandlers()

	return &a
}

func (a *App) registerHandlers() {
	a.dg.AddHandler(a.handleGuildCreate)
	a.dg.AddHandler(a.handleMessageCreateActivityRecord)
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
