package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"context"

	"net/http"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

const (
	secretKey = "102703av0grv8n4l"
)

func CookieHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userUUID := uuid.New().String()

		if cookieUserUUID, err := r.Cookie(model.CookieUserUUIDName); err != nil {
			// userUUID cookie
			cookieUserUUID := newCookie(model.CookieUserUUIDName, userUUID)
			http.SetCookie(w, cookieUserUUID)

			// signature cookie
			cookieSignature := newCookie(model.CookieSignatureName, stringToHmacSha256(userUUID))
			http.SetCookie(w, cookieSignature)
		} else {
			userUUID = cookieUserUUID.Value
			cookieSignature, _ := r.Cookie(model.CookieSignatureName)
			signature := cookieSignature.Value

			if strings.Compare(stringToHmacSha256(userUUID), signature) != 0 {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), model.CookieContextName, userUUID)))
	})
}

func stringToHmacSha256(data string) string {
	h := hmac.New(sha256.New, []byte(secretKey))

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
