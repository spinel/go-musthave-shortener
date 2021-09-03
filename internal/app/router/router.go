package router

import (
	"github.com/gorilla/mux"
	"github.com/spinel/go-musthave-shortener/internal/app/handler"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

// Router for an app.
func NewRouter(repo repository.Repositorer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.NewCreateEntityHandler(repo))
	r.HandleFunc("/{id:[0-9a-z]+}", handler.NewGetEntityHandler(repo))
	r.HandleFunc("/api/shorten", handler.NewCreateJSONEntityHandler(repo))

	return r
}
