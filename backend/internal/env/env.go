package env

import (
	"log/slog"
	"os"
)

func MustGet(envVar string) string {
	r := os.Getenv(envVar)

	if r == "" {
		slog.Error("env variable get", envVar, "unset")
		os.Exit(1)
	}

	return r
}
