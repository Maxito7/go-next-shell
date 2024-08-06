package main

import (
	"fmt"
	"ingesoft/backend/config"
	"ingesoft/backend/internal/auth"
	"ingesoft/backend/internal/route"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	e := echo.New()
	auth.InitAuth(e, cfg)
	route.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
	// e.Logger.Fatal(e.Start(":8080"))
}
