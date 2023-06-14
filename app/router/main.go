package router

import (
	"app/router/user"

	"github.com/labstack/echo/v4"
)

func Register(router *echo.Echo) {
	user.Register(router)
}
