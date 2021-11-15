package main

import (
	"goproject/echotest/controllers"
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

	e.GET("/users", controllers.GetUser)

	e.POST("/users", controllers.InsertUser)

	e.DELETE("/users/:id",controllers.DeleteUser)

	e.PUT("/users",controllers.EditUser)
	e.Logger.Fatal(e.Start(":8080"))
}
