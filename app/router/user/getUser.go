package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var usersData = []User{
	{
		Name: "John",
		Age:  20,
	},

	{
		Name: "Jane",
		Age:  30,
	},
}

func getUser(context echo.Context) error {
	return context.JSON(http.StatusOK, usersData)
}
