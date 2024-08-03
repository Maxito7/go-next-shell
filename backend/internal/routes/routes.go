package routes

import (
	"ingesoft/backend/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
	e.GET("/usersdata", handlers.GetData)
	e.GET("/auth/{provider}", handlers.ProviderAuth)
	e.GET("/auth/{provider}/callback", handlers.ProviderCallback)
	e.GET("/logout/{provider}", handlers.ProviderLogout)
}
