package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"test/internal/database"
	"test/internal/logger"
)

type Server struct {
	port   int
	logger *logger.Logger
	db     database.Service
}

func NewServer(logger *logger.Logger) (*http.Server, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db, err := database.New()
	if err != nil {
		return nil, fmt.Errorf("error creating new server: %w", err)
	}
	NewServer := &Server{port, logger, db}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
