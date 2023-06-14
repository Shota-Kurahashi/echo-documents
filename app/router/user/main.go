package user

import (
	"github.com/labstack/echo/v4"
)

func Register(router *echo.Echo) {
	router.GET("/user", getUser)
}
