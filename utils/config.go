package utils

import "github.com/caarlos0/env/v11"

type Config struct {
	BotToken string `env:"BOT_TOKEN,required,notEmpty"`
}

func MustLoadConfig() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
