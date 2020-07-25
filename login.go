package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func newUser() string {
	newUserUUID := uuid.NewV4()
	return newUserUUID.String()
}

func mainLoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if password != "r3jp0i" {
		return echo.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	newUser := newUser()
	claims["name"] = username
	claims["uuid"] = newUser
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
