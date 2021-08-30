package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/pkg"
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/repository"
)

const Host = "http://localhost:8080"

// CreateEntityHandler - save entity in the store.
func NewCreateEntityHandler(repo *repository.Store) func(w http.ResponseWriter, r *http.Request) {
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

		entity := model.Entity{
			URL: url,
		}
		var code string
		for {

			code, err = pkg.NewGeneratedString()
			if err != nil {
				http.Error(w, "save entity error", http.StatusInternalServerError)
				return
			}

			if !repo.Entity.IncludesCode(string(code)) {
				break
			}
		}

		err = repo.Entity.SaveEntity(code, entity)
		if err != nil {
			http.Error(w, "entity exists", http.StatusInternalServerError)

			return
		}

		result := fmt.Sprintf("%s/%s", Host, code)
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(result))
	}
}

// GetEntityHandler retrive entity from store by id.
func NewGetEntityHandler(repo *repository.Store) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")

		if len(pathSplit) != 2 {
			http.Error(w, "no id", http.StatusBadRequest)

			return
		}
		id := pathSplit[1]

		entity, err := repo.Entity.GetEntityBy(id)
		if err != nil {
			http.Error(w, "get entity error", http.StatusInternalServerError)
			return
		}

		if entity.URL == "" {
			http.Error(w, "not found", http.StatusNotFound)

			return
		}

		http.Redirect(w, r, entity.URL, http.StatusTemporaryRedirect)
	}
}
