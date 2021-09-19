package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const cookieName = "userUUId"
const cookieExpireDays = 30
const cookieContextName = "userUUId"

func CookieHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userUUId := uuid.New().String()

		if cookie, err := r.Cookie(cookieName); err != nil {
			cookie := &http.Cookie{
				Name:    cookieName,
				Value:   userUUId,
				Path:    "/",
				Expires: time.Now().Add(cookieExpireDays * 24 * time.Hour),
			}
			http.SetCookie(w, cookie)
		} else {
			userUUId = cookie.Value
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), cookieContextName, userUUId)))
	})
}
