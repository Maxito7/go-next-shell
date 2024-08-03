package auth

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func InitAuth(e *echo.Echo) {
	googleClientID := "161589305451-d8f7vqjvjh21dhuo7cqqd5as7vohpmb3.apps.googleusercontent.com"
	googleClientSecret := "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"
	callbackURL := "http://localhost:1323/auth/google/callback"

	store := sessions.NewCookieStore([]byte(key))

	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	e.Use(session.Middleware(store))

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, callbackURL),
	)
}
