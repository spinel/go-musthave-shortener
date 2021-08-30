package router

import (
	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/handler"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository"
)

// Router for an app.
func Router(repo *repository.Store) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.NewCreateEntityHandler(repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.NewGetEntityHandler(repo))

	return r
}
