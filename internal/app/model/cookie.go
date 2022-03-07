package model

import "github.com/spinel/go-musthave-shortener/internal/app/pkg"

const (
	CookieUserUUIDName  = "userUUId"
	CookieSignatureName = "signature"
	CookieExpireDays    = 30
	CookieContextName   = pkg.ContextKey("userUUID")
)
