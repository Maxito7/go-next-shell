package handlers

import (
	"context"
	"ingesoft/backend/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func GetData(c echo.Context) error {
	data, err := services.FetchData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, data)
}

func ProviderAuth(c echo.Context) error {
	/*
		provider := c.Param("provider")
		if provider != "" {
			return c.String(http.StatusBadRequest, "Provider not specified")
		}

		q := c.Request().URL.Query()
		q.Add("provider", c.Param("provider"))
		c.Request().URL.RawQuery = q.Encode()

		req := c.Request()
		res := c.Response().Writer
		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		// Hacer algo con el usuario, crear una sesion o algo xd
		log.Printf("User: %#v", user)
		gothic.BeginAuthHandler(res, req)

		return c.JSON(http.StatusOK, user)
	*/

	ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request().WithContext(ctx)); err == nil {
		return c.JSON(http.StatusOK, gothUser)
	}

	gothic.BeginAuthHandler(c.Response(), c.Request().WithContext(ctx))

	return nil
}

func ProviderCallback(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request().WithContext(ctx))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func ProviderLogout(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

	gothic.Logout(c.Response(), c.Request().WithContext(ctx))
	c.Response().Header().Set("Location", "/")
	c.Response().WriteHeader(http.StatusTemporaryRedirect)
	return nil
}
