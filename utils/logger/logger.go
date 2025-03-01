package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Koobson/kotbot/utils"
)

var logFileLock sync.Mutex

type kotbotLoggerHandler struct {
	level      slog.Level
	logDirPath string
}

func New(level slog.Level, logDirPath string) *kotbotLoggerHandler {
	return &kotbotLoggerHandler{
		level:      level,
		logDirPath: logDirPath,
	}
}

func (h *kotbotLoggerHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.level <= level
}

func (h *kotbotLoggerHandler) Handle(ctx context.Context, r slog.Record) error {
	messageBuilder := strings.Builder{}

	messageBuilder.WriteString(r.Time.Format("2006-01-02 15:04:05.000"))
	messageBuilder.WriteString(" ")
	messageBuilder.WriteString(r.Level.String())
	messageBuilder.WriteString(" ")
	messageBuilder.WriteString(r.Message)
	messageBuilder.WriteString(" ")
	r.Attrs(func(a slog.Attr) bool {
		if !a.Equal(slog.Attr{}) {
			messageBuilder.WriteString(a.Key)
			messageBuilder.WriteString("=")
			messageBuilder.WriteString(a.Value.String())
			messageBuilder.WriteString(" ")
		}
		return true
	})

	messageBuilder.WriteString("\n")
	fmt.Print(messageBuilder.String())
	err := h.addLogToFile(messageBuilder.String())
	if err != nil {
		panic(err)
	}
	return nil
}

func (h *kotbotLoggerHandler) addLogToFile(log string) error {
	curTime := time.Now()
	filename := fmt.Sprintf("log%d-%02d-%02d.log", curTime.Year(), curTime.Month(), curTime.Day())
	logPath := filepath.Clean(h.logDirPath)

	logFileLock.Lock()
	defer logFileLock.Unlock()

	filePathExists, err := utils.PathExists(logPath)
	if err != nil {
		return err
	}

	if !filePathExists {
		return fmt.Errorf("path to log logs: %s does not exist", logPath)
	}

	f, err := os.OpenFile(logPath+"/"+filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(log); err != nil {
		return err
	}

	return nil
}

func (h *kotbotLoggerHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("Unreachable")
}

func (h *kotbotLoggerHandler) WithGroup(name string) slog.Handler {
	panic("Unreachable")
}
