package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Koobson/kotbot/app"
	"github.com/Koobson/kotbot/storage"
	"github.com/Koobson/kotbot/utils"
	"github.com/Koobson/kotbot/utils/logger"
)

//TODO robić folder na logi jak go nie ma
//TODO robić folder na db jak go nie ma

func main() {
	cfg := utils.MustLoadConfig()
	slog.SetDefault(slog.New(logger.New(slog.LevelInfo, cfg.LogPath)))
	slog.Info("App Start", logger.BotVersion(cfg.CommitHash))

	storage := storage.New(cfg.DBPath + "/db.sqlite")
	defer storage.Close()

	botApp := app.New(cfg.BotToken, storage)

	_ = botApp

	botApp.Start()
	defer botApp.Stop()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
