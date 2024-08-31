package server

import (
	"context"
	"fmt"
	"go-project-setup/internal/logger"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

func NewServer() *Server {
	addr := getEnv("SERVER_ADDR", "8080")
	readTimeout := getEnvAsDuration("READ_TIMEOUT", 5*time.Second)
	writeTimeout := getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second)
	idleTimeout := getEnvAsDuration("IDLE_TIMEOUT", 15*time.Second)

	srv := &Server{
		router: chi.NewRouter(),
	}

	srv.registerRoutes()
	srv.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%s", addr),
		Handler:      srv.router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	return srv
}

func (s *Server) Start() error {
	logger.Log.Info("starting HTTP server", "address", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.Log.Info("shutting down HTTP server")
	return s.httpServer.Shutdown(ctx)
}

// getEnv reads an environment variable or returns a default value if not set
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsDuration reads an environment variable and converts it to a time.Duration
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		logger.Log.Error("invalid duration value", "key", key, "value", valueStr, "error", err)
		return defaultValue
	}
	return time.Duration(value) * time.Second
}
