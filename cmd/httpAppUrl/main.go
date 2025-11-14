package main

import (
	"log/slog"
	"os"
	"status-links/internal/app"
	"status-links/internal/config"
)

func main() {
	cfg := config.MustLoad()
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	app := app.NewApp(cfg)
	app.Run()
}
