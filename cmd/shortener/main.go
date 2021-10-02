package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/handler/middleware"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
	"github.com/spinel/go-musthave-shortener/internal/app/router"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	repo, err := repository.NewStorage(cfg)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr: cfg.ServerAddress,
		Handler: middleware.CookieHandle(cfg,
			middleware.GzipHandle(
				router.NewRouter(cfg, repo.EntityPg),
			),
		),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)

	// Set os signals to the chan
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
