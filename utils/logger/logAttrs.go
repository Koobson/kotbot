package logger

import "log/slog"

const (
	KeyError           = "error"
	KeyGuildID         = "guildID"
	KeyUserID          = "userID"
	KeyKickedUserCount = "kickedUserCount"
	KeyGuildCount      = "guildCount"
)

func Err(err error) slog.Attr {
	return slog.String(KeyError, err.Error())
}

func GuildID(guildID string) slog.Attr {
	return slog.String(KeyGuildID, guildID)
}

func UserID(userID string) slog.Attr {
	return slog.String(KeyUserID, userID)
}

func GuildCount(guildCount int) slog.Attr {
	return slog.Int(KeyGuildCount, guildCount)
}

func KickedUserCount(kickedUserCount int) slog.Attr {
	return slog.Int(KeyKickedUserCount, kickedUserCount)
}
