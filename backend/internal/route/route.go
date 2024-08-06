package route

import (
	"ingesoft/backend/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handler.HelloWorld)
	e.GET("/auth/{provider}", handler.ProviderAuth)
	e.GET("/auth/{provider}/callback", handler.ProviderCallback)
	e.GET("/logout/{provider}", handler.ProviderLogout)
}
