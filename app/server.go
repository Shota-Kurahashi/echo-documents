package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(context echo.Context) error {
	fmt.Println("Hello, World!")
	return context.String(http.StatusOK, "Hello, World!")
}

func main() {
	router := echo.New()

	router.GET("/", hello)

	if err := router.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
