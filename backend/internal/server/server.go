package server

import (
	"ingesoft/backend/config"
	"ingesoft/backend/internal/auth"
	"ingesoft/backend/internal/middleware"
	"ingesoft/backend/internal/routes"

	"github.com/labstack/echo/v4"
)

func Start() {
	// config.LoadConfig()
	e := echo.New()
	middleware.Init(e)
	routes.RegisterRoutes(e) // Inicializamos todas las rutas
	auth.InitAuth(e)

	port := config.GetEnv("SERVER_PORT", "1323")
	e.Logger.Fatal(e.Start(":" + port))
}
