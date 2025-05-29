package app

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"quotes_api/internal/app/config"
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
		fmt.Println("REST API running on " + cfg.RestPort)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("REST API quotes server error: ", err.Error())
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
		fmt.Printf("Shutdown with error: %v", err)
	}
	fmt.Println("Shutdown complete...")
}
