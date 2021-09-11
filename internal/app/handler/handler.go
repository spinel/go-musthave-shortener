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

const Host = "http://localhost:8080"

// CreateEntityHandler - save entity to the store handler.
func NewCreateEntityHandler(cfg *config.Config, repo repository.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "wrong body", http.StatusBadRequest)

			return
		}
		url := string(body)
		if url == "" {
			http.Error(w, "no body", http.StatusBadRequest)

			return
		}

		code, err := repo.GetCode(url)
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

// NewCreateJSONEntityHandler - API JSON version, save entity to the store handler.
// Get JSON in body, return Result as JSON.
func NewCreateJSONEntityHandler(cfg *config.Config, repo repository.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

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

		code, err := repo.GetCode(entity.URL)
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
