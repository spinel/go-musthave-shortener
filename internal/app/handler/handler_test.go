package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

const (
	testURL      = "https://yandex.ru/"
	testUserUUID = "7c03d351-d0e4-41d0-a837-6df16ced19d4"
)

func TestNewGetUrlHandler(t *testing.T) {
	const testCode = "testtest"
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		path string
		want want
	}{
		{
			name: "#5 get request test",
			path: testCode,
			want: want{
				code:        http.StatusTemporaryRedirect,
				contentType: "application/text",
			},
		},
		{
			name: "#6 get request test undefined code",
			path: "_",
			want: want{
				code:        http.StatusNotFound,
				contentType: "application/text",
			},
		},
	}

	cfg := config.NewConfig()
	repoStorage, err := repository.NewStorage(cfg)
	if err != nil {
		panic(err)
	}
	userUUID, _ := uuid.Parse(testUserUUID)

	_, err = repoStorage.EntityPg.CreateURL(&model.Entity{
		URL:      testURL,
		Code:     testCode,
		UserUUID: userUUID,
	})

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", fmt.Sprintf("/%s", tc.path), nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(NewGetUrlHandler(repoStorage.EntityPg))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tc.want.code, res.StatusCode)
		})
	}
}

func TestNewCreateUrlHandler(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}

	tests := []struct {
		name    string
		payload string
		want    want
	}{
		{
			name:    "#1 post request test good payload",
			payload: testURL,
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:    "#2 post request test empty payload",
			payload: "",
			want: want{
				code: http.StatusBadRequest,
				//				response:    "",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}
	ctx := context.Background()

	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	repoStorage, err := repository.NewStorage(cfg)
	if err != nil {
		panic(err)
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", "/", strings.NewReader(tc.payload))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(NewCreateUrlHandler(cfg, repoStorage.EntityPg))
			h.ServeHTTP(w, request.WithContext(context.WithValue(ctx, model.CookieContextName, testUserUUID)))

			res := w.Result()
			defer res.Body.Close()

			//status code
			assert.EqualValues(t, tc.want.code, res.StatusCode)

			//content-type
			assert.EqualValues(t, res.Header.Get("Content-Type"), tc.want.contentType)
		})
	}
}

func TestNewCreateJSONEntityHandler(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	entity := model.Entity{
		URL: testURL,
	}
	testBody, _ := json.Marshal(entity)
	tests := []struct {
		name    string
		payload string
		want    want
	}{
		{
			name:    "#3 post request test good payload",
			payload: string(testBody),
			want: want{
				code:        http.StatusCreated,
				contentType: "application/json; charset=utf-8",
			},
		},
		{
			name:    "#4 post request test empty payload",
			payload: "",
			want: want{
				code:        http.StatusBadRequest,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	ctx := context.Background()

	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	repoStorage, err := repository.NewStorage(cfg)
	if err != nil {
		panic(err)
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", "/", strings.NewReader(tc.payload))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(NewCreateJsonUrlHandler(cfg, repoStorage.EntityPg))
			h.ServeHTTP(w, request.WithContext(context.WithValue(ctx, model.CookieContextName, testUserUUID)))
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tc.want.code, res.StatusCode)

			//content-type
			assert.EqualValues(t, tc.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
