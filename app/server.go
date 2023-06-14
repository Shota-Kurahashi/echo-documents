package main

import (
	"app/router"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(context echo.Context) error {

	return context.String(http.StatusOK, "Hello, World!")
}

func main() {
	server := echo.New()

	server.GET("/", hello)
	router.Register(server)

	if err := server.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
