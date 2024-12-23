package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"github.com/benjaspet/generate-s25-software-challenge/src/internal/server"
)

func main() {
    
    app := server.New()

    go func() {
        if err := app.Listen(":8080"); err != nil {
            log.Fatalf("Failed to start server: %v", err)
        }
    }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	slog.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}

	slog.Info("Server shutdown")
}