package main

import (
	_ "goproject/echotest/controllers"
	_ "goproject/echotest/models"
	"net/http"

	"github.com/labstack/echo" // set GOPATH module
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", editUser )
	e.Logger.Fatal(e.Start(":8080"))
}
