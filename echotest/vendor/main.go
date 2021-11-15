package main

import (
	_ "goproject/echotest/controllers"
	"net/http"

	"github.com/labstack/echo" // set GOPATH module
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// e.GET("/user", func(c echo.Context) error {
	// 	return c.JASON(http.StatusOK, "/user")
	// })
	e.POST("/users", createUser)
	e.GET("/users", getUserModel)
	e.Logger.Fatal(e.Start(":8080"))
}
