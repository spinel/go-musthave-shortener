package router

import (
	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/handler"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

// Router for an app.
func NewRouter(cfg *config.Config, repo repository.URLShortener) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.NewCreateEntityHandler(cfg, repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.NewGetEntityHandler(repo))
	r.HandleFunc("/user/urls", handler.NewGetUserURLSHandler(repo))
	r.HandleFunc("/api/shorten", handler.NewCreateJSONEntityHandler(cfg, repo))

	return r
}
