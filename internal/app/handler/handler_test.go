package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

const testURL = "https://yandex.ru/"

func TestNewCreateEntityHandler(t *testing.T) {
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
	repo, err := repository.New()
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", "/", strings.NewReader(tt.payload))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(NewCreateEntityHandler(repo))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)

			//content-type
			assert.EqualValues(t, res.Header.Get("Content-Type"), tt.want.contentType)
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
	repo, err := repository.New()
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", "/", strings.NewReader(tt.payload))
			w := httptest.NewRecorder()
			h := http.HandlerFunc(NewCreateJSONEntityHandler(repo))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, tt.want.code, res.StatusCode)

			//content-type
			assert.EqualValues(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
func TestNewGetEntityGetEntityHandler(t *testing.T) {
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
			name: "#6 get request test",
			path: "_",
			want: want{
				code:        http.StatusNotFound,
				contentType: "application/text",
			},
		},
	}
	repo, err := repository.New()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Entity.SaveEntity(testCode, model.Entity{URL: testURL})
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", fmt.Sprintf("/%s", tt.path), nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(NewGetEntityHandler(repo))
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			//status code
			assert.EqualValues(t, res.StatusCode, tt.want.code)
		})
	}
}
