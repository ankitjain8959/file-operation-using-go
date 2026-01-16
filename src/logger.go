package src

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

// initialize the logger once
func InitLogger() {
	Log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)
}
