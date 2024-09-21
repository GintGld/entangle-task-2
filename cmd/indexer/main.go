package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"task2/internal/app"
	"task2/internal/config"
)

func main() {
	// Read config.
	cfg := config.LoadConfig()

	// Setup logger.
	log := setupLogger()

	// Create application.
	cliApp := app.New(
		log,
		cfg.RPC,
		cfg.Out,
	)

	// Run indexer
	cliApp.Indexer.Run(cfg.InitNum)

	// Graceful shutdown.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	<-stop

	// Stop application.
	cliApp.Indexer.Close()
	log.Info("Gracefully stopped")
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)
}
