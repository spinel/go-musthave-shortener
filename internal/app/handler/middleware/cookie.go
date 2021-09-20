package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

func CookieHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userUUId := uuid.New().String()

		if cookie, err := r.Cookie(model.CookieName); err != nil {
			cookie := &http.Cookie{
				Name:    model.CookieName,
				Value:   userUUId,
				Path:    "/",
				Expires: time.Now().Add(model.CookieExpireDays * 24 * time.Hour),
			}
			http.SetCookie(w, cookie)
		} else {
			userUUId = cookie.Value
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.CookieContextName, userUUId)))
	})
}
