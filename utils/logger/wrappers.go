package logger

import (
	"context"
	"log/slog"
)

func LogCommand(message string, err error, guildID string, userID string, attrs ...slog.Attr) {
	logAttrs := []any{}
	level := slog.LevelInfo
	if err != nil {
		level = slog.LevelError
		logAttrs = append(logAttrs, Err(err))
	}

	logAttrs = append(logAttrs, GuildID(guildID), UserID(userID))
	for attr := range attrs {
		logAttrs = append(logAttrs, attr)
	}

	slog.Log(context.Background(), level, message, logAttrs...)
}

func Log(message string, err error, attrs ...slog.Attr) {
	logAttrs := []any{}
	level := slog.LevelInfo
	if err != nil {
		level = slog.LevelError
		logAttrs = append(logAttrs, Err(err))
	}

	for attr := range attrs {
		logAttrs = append(logAttrs, attr)
	}

	slog.Log(context.Background(), level, message, logAttrs...)
}
