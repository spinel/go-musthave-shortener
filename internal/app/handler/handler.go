package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

// CreateEntityHandler - save entity to the store handler.
func NewCreateEntityHandler(cfg *config.Config, repo repository.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		ctx := r.Context()

		if err != nil {
			http.Error(w, "wrong body", http.StatusBadRequest)

			return
		}
		url := string(body)
		if url == "" {
			http.Error(w, "no body", http.StatusBadRequest)

			return
		}

		code, err := repo.GetCode(ctx, url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result := fmt.Sprintf("%s/%s", cfg.BaseURL, code)
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)

		w.Write([]byte(result))
	}
}

// GetEntityHandler retrive entity from store by id handler.
func NewGetEntityHandler(repo repository.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")

		if len(pathSplit) != 2 {
			http.Error(w, "no id", http.StatusBadRequest)

			return
		}
		id := pathSplit[1]

		entity, err := repo.GetEntityBy(id)
		if err != nil {
			http.Error(w, "entity not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, entity.URL, http.StatusTemporaryRedirect)
	}
}

// NewGetUserURLSHandler retrive current user urls
func NewGetUserURLSHandler(cfg *config.Config, repo repository.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		urlMappingPool := repo.GetByUser(ctx, cfg)

		if len(urlMappingPool) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Add("Content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(urlMappingPool)
	}
}

// NewCreateJSONEntityHandler - API JSON version, save entity to the store handler.
// Get JSON in body, return Result as JSON.
func NewCreateJSONEntityHandler(cfg *config.Config, repo repository.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		ctx := r.Context()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}
		entity := model.Entity{}
		err = json.Unmarshal(body, &entity)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		code, err := repo.GetCode(ctx, entity.URL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result := model.Result{
			URL: fmt.Sprintf("%s/%s", cfg.BaseURL, code),
		}
		w.Header().Add("Content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(result)
	}
}

// NewPingHandler for check pg db connection
func NewPingHandler(repo repository.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if !repo.Ping() {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
