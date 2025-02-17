package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Koobson/kotbot/app"
	"github.com/Koobson/kotbot/utils"
)

func main() {
	cfg := utils.MustLoadConfig()

	botApp := app.New(cfg.BotToken)

	_ = botApp

	botApp.Start()
	defer botApp.Stop()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
