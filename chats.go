package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func chatsHelloHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	uuid := claims["uuid"].(string)
	return c.JSON(http.StatusOK, map[string]string{
		"name": name,
		"uuid": uuid,
	})
}
