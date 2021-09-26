package router

import (
	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/handler"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

// Router for an app.
func NewRouter(cfg *config.Config, repo repository.UrlStorer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.NewCreateUrlHandler(cfg, repo))
	r.HandleFunc("/api/shorten", handler.NewCreateJsonUrlHandler(cfg, repo))
	r.HandleFunc("/ping", handler.NewPingHandler(repo))
	r.HandleFunc("/user/urls", handler.NewGetUserUrlsHandler(cfg, repo))
	r.HandleFunc("/api/shorten/batch", handler.NewCreateBatchHandler(cfg, repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.NewGetUrlHandler(repo))

	return r
}
