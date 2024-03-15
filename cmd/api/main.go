package main

import (
	"log"
	"log/slog"
	"os"
	"test/internal/logger"
	"test/internal/server"
)

func main() {
	logger, err := logger.NewLogger(logger.LoggerConfig{
		LoggerType: "text",
		Level:      slog.LevelDebug,
		AddSource:  true,
	})
	if err != nil {
		log.Fatalf("cannot create logger: %s. Cannot initialise...", err)
	}

	server, err := server.NewServer(logger)
	if err != nil {
		logger.Error("cannot create server", slog.String("error", err.Error()))
		logger.Info("Exiting...")
		os.Exit(1)
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Error("cannot start server", slog.String("error", err.Error()))
		logger.Info("Exiting...")
		os.Exit(1)
	}
}
