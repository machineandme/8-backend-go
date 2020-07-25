package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var htmlDocument string

const secret = "Kwfmiojp3__r2@Rmnwfi--eiF@# pioejgv@MNpv i 2i3nrF@#N@# Nwnfegi@F#NI f3@"

func init() {
	b, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Print(err)
	}
	htmlDocument = string(b)
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://e0044380f751471aaa22920b66a4e18c@o423712.ingest.sentry.io/5359412",
		Release:     "0.0.1",
		DebugWriter: ioutil.Discard,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	e := echo.New()
	e.Use(sentryecho.New(sentryecho.Options{
		Repanic:         false,
		WaitForDelivery: false,
		Timeout:         time.Duration(200 * time.Millisecond),
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, htmlDocument)
	})

	e.POST("/login", mainLoginHandler)

	r := e.Group("/chats")
	r.Use(middleware.JWT([]byte(secret)))
	r.GET("/chats", chatsHelloHandler)

	// Start server
	e.Logger.Fatal(e.Start("localhost:1323"))
}
