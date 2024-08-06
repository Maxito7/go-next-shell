package auth

import (
	"ingesoft/backend/config"

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

func InitAuth(e *echo.Echo, cfg *config.Config) {
	googleClientID := cfg.GoogleAuth.ClientID
	googleClientSecret := cfg.GoogleAuth.ClientSecret
	callbackURL := cfg.GoogleAuth.CallbackURL

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
