package router

import (
	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/handler"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

// Router for an app.
func NewRouter(cfg *config.Config, repo repository.URLStorer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.NewCreateURLHandler(cfg, repo))
	r.HandleFunc("/api/shorten", handler.NewCreateJsonURLHandler(cfg, repo))
	r.HandleFunc("/ping", handler.NewPingHandler(repo))
	r.HandleFunc("/user/urls", handler.NewGetUserURLsHandler(cfg, repo))
	r.HandleFunc("/api/shorten/batch", handler.NewCreateBatchHandler(cfg, repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.NewGetURLHandler(repo))

	return r
}
