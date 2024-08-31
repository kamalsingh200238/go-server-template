package main

import (
	"context"
	"go-project-setup/internal/logger"
	"go-project-setup/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// init logger
	logger.Init()

	// Initialize the server
	srv := server.NewServer()

	// Start the server in a goroutine
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error("error in starting the server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Error("error in shutting down server", "error", err)
	}

	logger.Log.Info("server exiting")
}
