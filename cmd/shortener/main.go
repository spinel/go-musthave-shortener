package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/web"
	"github.com/spinel/go-musthave-shortener/internal/app/router"
)

func main() {
	// Init store
	db := make(map[string]model.Entity)

	// Entity interface
	entityRepo := web.NewEntityRepo(db)

	server := &http.Server{Addr: ":8080", Handler: router.NewRouter(entityRepo)}

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
	// Wait for ListenAndServe goroutine to close.
}
