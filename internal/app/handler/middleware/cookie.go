package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"context"

	"net/http"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

type userUUID string

func toUserUUID(s string) userUUID {
	return userUUID(s)
}

func CookieHandle(cfg config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userUUID userUUID

		if cookieUserUUID, err := r.Cookie(model.CookieUserUUIDName); err != nil {
			userUUID := uuid.New().String()
			// userUUID cookie
			cookieUserUUID := newCookie(model.CookieUserUUIDName, userUUID)
			http.SetCookie(w, cookieUserUUID)

			// signature cookie
			cookieSignature := newCookie(model.CookieSignatureName, stringToHmacSha256(cfg, userUUID))
			http.SetCookie(w, cookieSignature)
		} else {
			userUUID = toUserUUID(cookieUserUUID.Value)
			cookieSignature, _ := r.Cookie(model.CookieSignatureName)
			signature := cookieSignature.Value

			if strings.Compare(stringToHmacSha256(cfg, string(userUUID)), signature) != 0 {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.CookieContextName, userUUID)))
	})
}

func stringToHmacSha256(cfg config.Config, data string) string {
	h := hmac.New(sha256.New, []byte(cfg.CookieSecretKey))

	// Write Data
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}

func newCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}
}
