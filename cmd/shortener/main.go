package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/handler/middleware"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
	"github.com/spinel/go-musthave-shortener/internal/app/router"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

//syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT
func main() {
	fmt.Printf("Build version: %s\n", pkg.CheckNA(buildVersion))
	fmt.Printf("Build date: %s\n", pkg.CheckNA(buildDate))
	fmt.Printf("Build commit: %s\n\n", pkg.CheckNA(buildCommit))

	ctx := context.Background()
	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		log.Fatal("config validation failed: ", err)
	}
	if cfg.Config != "" {
		cfg, _ = pkg.ReadJsonFile(cfg.Config)
	}

	repo, err := repository.NewStorage(cfg)
	if err != nil {
		log.Fatal("storage init failed:", err)
	}

	server := &http.Server{
		Addr: cfg.ServerAddress,
		Handler: middleware.CookieHandle(cfg,
			middleware.GzipHandle(
				router.NewRouter(cfg, repo.EntityPg),
			),
		),
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if cfg.EnableHttps {
			if err := server.ListenAndServeTLS("", ""); err != nil {
				log.Fatal("tls server start failed:", err)
			}
		} else {
			if err := server.ListenAndServe(); err != nil {
				log.Fatal("server start failed:", err)
			}
		}
		wg.Done()
	}()
	wg.Wait()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)

	// Set os signals to the chan
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("graceful shutdown failed: ", err)
	}
}
