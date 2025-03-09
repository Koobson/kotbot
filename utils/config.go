package utils

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	CommitHash string `env:"COMMIT_HASH,required,notEmpty"`
	BotToken   string `env:"BOT_TOKEN,required,notEmpty"`
	DBPath     string `env:"DB_PATH,required,notEmpty"`
	LogPath    string `env:"LOG_PATH,required,notEmpty"`
}

func MustLoadConfig() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
