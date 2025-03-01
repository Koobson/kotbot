package logger

import (
	"context"
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.String(KeyError, err.Error())
}

func LogCommand(commandName string, err error, guildID string, userID string, attrs ...slog.Attr) {
	logAttrs := []any{}
	level := slog.LevelInfo
	if err != nil {
		level = slog.LevelError
		logAttrs = append(logAttrs, Err(err))
	}

	logAttrs = append(logAttrs, slog.String(GuildID, guildID), slog.String(UserID, userID))
	for attr := range attrs {
		logAttrs = append(logAttrs, attr)
	}

	slog.Log(context.Background(), level, commandName, logAttrs...)
}
