package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

// NewPingHandler for check pg db connection
func NewPingHandler(repo repository.URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if !repo.Ping() {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// NewCreateURLHandler - save new entity handler.
func NewCreateURLHandler(cfg *config.Config, repo repository.URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
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

		urlCode, err := pkg.NewGeneratedString()
		if err != nil {
			http.Error(w, "code generate error", http.StatusInternalServerError)
		}

		userUUID := getUserUUIDFromCtx(ctx)

		entity := &model.Entity{
			Code:     urlCode,
			URL:      url,
			UserUUID: userUUID,
		}

		existEntity, err := repo.CreateURL(entity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Add("Content-type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)

		if existEntity != nil {
			urlCode = existEntity.Code
			w.WriteHeader(http.StatusConflict)
		}

		result := pkg.FormatLocalURL(cfg.BaseURL, urlCode)

		w.Write([]byte(result))
	}
}

// NewCreateJsonURLHandler - API JSON version, save entity to the store handler.
// Get JSON in body, return Result as JSON.
func NewCreateJsonURLHandler(cfg *config.Config, repo repository.URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		ctx := r.Context()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		urlCode, err := pkg.NewGeneratedString()
		if err != nil {
			http.Error(w, "code generate error", http.StatusInternalServerError)
		}

		entity := &model.Entity{}
		err = json.Unmarshal(body, &entity)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		userUUID := getUserUUIDFromCtx(ctx)
		entity.UserUUID = userUUID
		entity.Code = urlCode

		existEntity, err := repo.CreateURL(entity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Add("Content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)

		if existEntity != nil {
			urlCode = existEntity.Code
			w.WriteHeader(http.StatusConflict)
		}

		result := model.Result{
			URL: pkg.FormatLocalURL(cfg.BaseURL, urlCode),
		}

		json.NewEncoder(w).Encode(result)
	}
}

// NewGetURLHandler retrive entity from store by code handler.
func NewGetURLHandler(repo repository.URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSplit := strings.Split(r.URL.Path, "/")

		if len(pathSplit) != 2 {
			http.Error(w, "no code", http.StatusBadRequest)

			return
		}
		urlCode := pathSplit[1]

		entity, err := repo.GetURL(urlCode)
		if entity == nil {
			http.Error(w, "entity not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, entity.URL, http.StatusTemporaryRedirect)
	}
}

// NewCreateBatchHandler - mass list of urls save.
func NewCreateBatchHandler(cfg *config.Config, repo repository.URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userUUID := getUserUUIDFromCtx(ctx)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "wrong body", http.StatusBadRequest)

			return
		}

		var batchURLs []*model.RequestBatchURLS

		err = json.Unmarshal(body, &batchURLs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var entities []model.Entity
		for _, batchURL := range batchURLs {
			urlCode, err := pkg.NewGeneratedString()
			if err != nil {
				http.Error(w, "code generate error", http.StatusInternalServerError)
			}

			entity := model.Entity{
				Code:     urlCode,
				URL:      batchURL.OriginalURL,
				UserUUID: userUUID,
			}
			entities = append(entities, entity)

			batchURL.ShortURL = pkg.FormatLocalURL(cfg.BaseURL, urlCode)
		}

		err = repo.SaveBatch(ctx, entities)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(batchURLs)
	}
}

// NewGetUserURLsHandler retrive current user urls
func NewGetUserURLsHandler(cfg *config.Config, repo repository.URLStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		userUUID := getUserUUIDFromCtx(ctx)

		entities := repo.GetByUser(ctx, userUUID)

		// convert Entity to URLMapping
		var urlMappingPool []model.URLMapping
		for _, entity := range entities {
			urlMappingPool = append(urlMappingPool, entity.ToURLMapping(cfg))
		}

		if len(urlMappingPool) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Add("Content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(urlMappingPool)
	}
}

func getUserUUIDFromCtx(ctx context.Context) uuid.UUID {
	userUUIDString := ctx.Value(model.CookieContextName).(string)
	userUUID, _ := uuid.Parse(userUUIDString)
	return userUUID
}
