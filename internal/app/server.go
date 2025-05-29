package app

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"quotes_api/internal/app/config"
	"quotes_api/internal/logging"
	"syscall"
	"time"
)

func RunServer() {
	cfg := config.MustLoad()
	router := setupRoutes()

	srv := http.Server{
		Addr:    cfg.RestPort,
		Handler: router,
	}

	go func() {
		logging.Logger.Info("REST API running on " + cfg.RestPort)
		if err := srv.ListenAndServe(); err != nil {
			logging.Logger.Error("Quotes REST error", "error", err)
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-shutdown.Done()

	// graceful shutdown
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logging.Logger.Error("Shutdown error", "error", err)
	} else {
		logging.Logger.Info("Shutdown complete...")
	}
}
