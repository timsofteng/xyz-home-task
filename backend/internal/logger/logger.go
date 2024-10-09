package logger

import (
	"log/slog"
	"os"
	"strings"
)

type Logger = *slog.Logger

func New(logLevelFlag string) Logger {
	logLevel := getLogLevel(logLevelFlag)

	options := &slog.HandlerOptions{
		Level:       logLevel,
		ReplaceAttr: replacer,
	}

	if logLevel == slog.LevelDebug {
		options.AddSource = true
	}

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, options))

	// slog.SetDefault(logger)
	logger.Info("Logger has been set", "logLevel", logLevel)

	return logger
}

// remove absolute path from source logs
func replacer(groups []string, a slog.Attr) slog.Attr {
	wd, _ := os.Getwd()
	if a.Key == slog.SourceKey {
		source := a.Value.Any().(*slog.Source)
		if file, ok := strings.CutPrefix(source.File, wd); ok {
			source.File = file
		}
	}
	return a

}

func getLogLevel(logLevelFlag string) slog.Level {
	logLevelMap := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	logLevel, ok := logLevelMap[logLevelFlag]

	if !ok {
		logLevel = slog.LevelInfo
	}

	return logLevel
}
